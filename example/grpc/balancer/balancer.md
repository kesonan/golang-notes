# Balancer
gRPC balancer


## 背景
接着上篇《gRPC 插件式编程之Resolver》，gRPC 将 `target` 解析为 `resolver.Target` 后，通过 `resolver.Builder.Build` 方法调用
`resolver.ClientConn.UpdateState(State) error` 方法，该方法做了哪些事情呢，我们本篇接着看源码往下走。

## UpdateState
UpdateState 的调用会调用 `grpc.ClientConn.updateResolverState` 方法，该方法主要做了如下工作：
- ServiceConfig 处理
- BalancerWrapper 创建
- 调用 `balancer.updateClientConnState` 方法 执行负载均衡逻辑更新

```go
func (cc *ClientConn) updateResolverState(s resolver.State, err error) error {
    ...
    cc.maybeApplyDefaultServiceConfig(s.Addresses)
    ...
    cc.applyServiceConfigAndBalancer(sc, configSelector, s.Addresses)
    ...
    // reference: balancer_conn_wrappers.go:164
    // bw.updateClientConnState -> ccBalancerWrapper.updateClientConnState
    bw.updateClientConnState(&balancer.ClientConnState{ResolverState: s, BalancerConfig: balCfg})
    ...
}
```
> ### 温馨提示
> 这里先以搞懂 gRPC 主流程思路为主，不扣太细节的东西，比如一些 `GRPCLB` 处理、error处理，ServiceConfigSelector 处理等可以查看源码。

`bw.updateClientConnState` 调用本质是 `ccBalancerWrapper.updateClientConnState`
而 `ccBalancerWrapper.updateClientConnState` 就做了一件事情，调用 `balancer.Balancer.UpdateClientConnState` 方法
```go
func (ccb *ccBalancerWrapper) updateClientConnState(ccs *balancer.ClientConnState) error {
    ccb.balancerMu.Lock()
    defer ccb.balancerMu.Unlock()
    return ccb.balancer.UpdateClientConnState(*ccs)
}
```

到这里，我们想看 `balancer` 源码逻辑有两种途径
- 自己实现的 `balancer.Balancer` 
- gRPC 提供的 `balancer`

为了阅读源码，我们先去阅读 gRPC 提供的几个 `balancer` 中的一个进行流程理解，后续再介绍如何自定义一个 `balancer`

## gRPC Balancer
gRPC 提供了几个负载均衡处理，如下：
- grpclb
- rls
- roundrobin
- weightroundrobin
- weighttarget

为了好理解，我们挑一个简单的负载均衡器 `roundrobin` 继续阅读。

负载均衡从哪里获取？通过前面 `cc.maybeApplyDefaultServiceConfig(s.Addresses)` 方法中的源码可知，`balancer.Balancer` 由 `balancer.Builder`
提供，我们看一下 `balancer.Builder` 接口
```go
// Builder creates a balancer.
type Builder interface {
    // Build creates a new balancer with the ClientConn.
    Build(cc ClientConn, opts BuildOptions) Balancer
    // Name returns the name of balancers built by this builder.
    // It will be used to pick balancers (for example in service config).
    Name() string
}
```

## roundrobin
roundrobin 是 gRPC 内置的负载均衡器，其和 `resolver` 一样都是通过插件式编程提供扩展，在源码中，我们可知，
roundrobin 在 `init` 函数中对 `balancer.Builder` 进行了注册，其中 `baseBuilder` 是 `balancer.Builder` 的实现，
上文我们得知， `balancer.Balancer` 由 `balancer.Builder.Build` 提供，通过 `baseBuilder.Build` 方法我们知道 gRPC 的
`balancer` 底层是由 `baseBalancer` 实现，部分源码如下：

roundrobin.go
```go
// newBuilder creates a new roundrobin balancer builder.
func newBuilder() balancer.Builder {
    return base.NewBalancerBuilder(Name, &rrPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
    balancer.Register(newBuilder())
}
```

balancer.go
```go
func (bb *baseBuilder) Build(cc balancer.ClientConn, opt balancer.BuildOptions) balancer.Balancer {
    bal := &baseBalancer{
        cc:            cc,
        pickerBuilder: bb.pickerBuilder,
    
        subConns: resolver.NewAddressMap(),
        scStates: make(map[balancer.SubConn]connectivity.State),
        csEvltr:  &balancer.ConnectivityStateEvaluator{},
        config:   bb.config,
    }
    bal.picker = NewErrPicker(balancer.ErrNoSubConnAvailable)
    return bal
}
```

沿着 `UpdateState` 环节最后一个方法 `ccb.balancer.UpdateClientConnState(*ccs)` 调用阅读，其实最终来到了
`baseBalancer.UpdateClientConnState` 方法，我们查看一下源码：

```go
func (b *baseBalancer) UpdateClientConnState(s balancer.ClientConnState) error {
    ...
    addrsSet := resolver.NewAddressMap()
    for _, a := range s.ResolverState.Addresses {
        addrsSet.Set(a, nil)
        if _, ok := b.subConns.Get(a); !ok {
            sc, err := b.cc.NewSubConn([]resolver.Address{a}, balancer.NewSubConnOptions{HealthCheckEnabled: b.config.HealthCheck})
            if err != nil {
                logger.Warningf("base.baseBalancer: failed to create new SubConn: %v", err)
                continue
            }
            b.subConns.Set(a, sc)
            b.scStates[sc] = connectivity.Idle
            b.csEvltr.RecordTransition(connectivity.Shutdown, connectivity.Idle)
            sc.Connect()
        }
    }
    for _, a := range b.subConns.Keys() {
        sci, _ := b.subConns.Get(a)
        sc := sci.(balancer.SubConn)
        if _, ok := addrsSet.Get(a); !ok {
            b.cc.RemoveSubConn(sc)
            b.subConns.Delete(a)
        }
    }
    if len(s.ResolverState.Addresses) == 0 {
        b.ResolverError(errors.New("produced zero addresses"))
        return balancer.ErrBadResolverState
    }
    return nil
}
```

从源码得知，该方法做了以下几件事：
- 对新的 endpoint `NewSubConn` 并且 `Connect`
- 移出旧的已经不存在的 `endpoint` 及其 `Conn` 信息

总的来说就是更新负载均衡器内可用的链接信息。

### balancer.ClientConn.NewSubConn
`balancer.ClientConn` 是一个接口，其代表 gRPC 的一个链接，而 `ccBalancerWrapper` 就为其实现类，先看看该接口的声明：
```go
type ClientConn interface {
    // NewSubConn 平衡器调用 NewSubConn 来创建一个新的SubConn，它不会阻塞并等待建立连接，
    // SubConn 的行为可以通过 NewSubConnOptions 来控制。
    NewSubConn([]resolver.Address, NewSubConnOptions) (SubConn, error)

    // RemoveSubConn 从ClientConn 中删除SubConn 。 SubConn将关闭。
    RemoveSubConn(SubConn)
    // UpdateAddresses 更新传入的SubConn 中使用的地址， gRPC 检查当前连接的地址是否仍在新列表中。 如果存在，将保持连接，
    // 否则，连接将正常关闭，并创建一个新连接。
    // 这将触发SubConn的状态转换。
    
    UpdateAddresses(SubConn, []resolver.Address)
    
    // UpdateState 通知 gRPC 平衡器的内部状态已更改。
    // gRPC 将更新ClientConn的连接状态，并在新的Picker上调用 Pick 来选择新的 SubConn。
    UpdateState(State)
    
    // 平衡器调用 ResolveNow 以通知 gRPC 进行名称解析。
    ResolveNow(resolver.ResolveNowOptions)
    
    // Target 返回此ClientConn的拨号目标。
    // 已弃用：改用BuildOptions中的 Target 字段
    Target() string
}
```

再看一下 `ccBalancerWrapper` 的创建:
```go
func newCCBalancerWrapper(cc *ClientConn, b balancer.Builder, bopts balancer.BuildOptions) *ccBalancerWrapper {
    ccb := &ccBalancerWrapper{
        cc:       cc,
        updateCh: buffer.NewUnbounded(),
        closed:   grpcsync.NewEvent(),
        done:     grpcsync.NewEvent(),
        subConns: make(map[*acBalancerWrapper]struct{}),
    }
    go ccb.watcher()
    ccb.balancer = b.Build(ccb, bopts)
    _, ccb.hasExitIdle = ccb.balancer.(balancer.ExitIdler)
    return ccb
}
```
> ### 注意
> 记住 `go ccb.watcher()` 这一行代码，后续还会回到这个方法来。


`baseBalancer.UpdateClientConnState` 中对新加入的 `endpoint` 进行 `NewSubConn` 和 `Connect` 处理，我们先来看看 `NewSubConn` 方法做了哪些事情，
来到 `ccBalancerWrapper.NewSubConn` 方法中:
```go
func (ccb *ccBalancerWrapper) NewSubConn(addrs []resolver.Address, opts balancer.NewSubConnOptions) (balancer.SubConn, error) {
	if len(addrs) <= 0 {
		return nil, fmt.Errorf("grpc: cannot create SubConn with empty address list")
	}
	ccb.mu.Lock()
	defer ccb.mu.Unlock()
	if ccb.subConns == nil {
		return nil, fmt.Errorf("grpc: ClientConn balancer wrapper was closed")
	}
	ac, err := ccb.cc.newAddrConn(addrs, opts)
	if err != nil {
		return nil, err
	}
	acbw := &acBalancerWrapper{ac: ac}
	acbw.ac.mu.Lock()
	ac.acbw = acbw
	acbw.ac.mu.Unlock()
	ccb.subConns[acbw] = struct{}{}
	return acbw, nil
}
```
从该方法可知，主要是通过 `gprc.ClientConn.newAddrConn` 创建一个 `addrConn` 对象，并且创建一个
`balancer.SubConn` 的实现类对象 `acBalancerWrapper`，将其加入到 `ccBalancerWrapper.subConns` 中进行管理。

> ### 说明
> 由此可知，`baseBalancer.UpdateClientConnState` 判断地址变更后的 address 是否为新加入的就由
> `ccBalancerWrapper.subConns` 来对比即可得知。

接着我们继续看看 `Connect` 做了什么事情，上面已经通过 `acBalancerWrapper` 创建了一个 `balancer.SubConn` 的实现对象，接着利用该对象进行了
`Connect` 方法调用，我们来到 `acBalancerWrapper.Connect()` 方法中：
```go
func (acbw *acBalancerWrapper) Connect() {
    acbw.mu.Lock()
    defer acbw.mu.Unlock()
    go acbw.ac.connect()
}
```

```go
func (ac *addrConn) connect() error {
    ac.mu.Lock()
    if ac.state == connectivity.Shutdown {
        ac.mu.Unlock()
        return errConnClosing
    }
    if ac.state != connectivity.Idle {
        ac.mu.Unlock()
        return nil
    }
    ac.updateConnectivityState(connectivity.Connecting, nil)
    ac.mu.Unlock()
    
    ac.resetTransport()
    return nil
}
```

`ac.updateConnectivityState` 更新链接状态，`ac.resetTransport` 主要工作内容就是从 `resolver.Address` 列表中按照去创建链接并同样调用 `ac.updateConnectivityState` 更新状态，具体源码可自行阅读，
我们接着 `ac.updateConnectivityState` 方法往下走，其实该方法调用了 `grpc.ClientConn.handleSubConnStateChange` 方法，最终又回到了 `ccBalancerWrapper.handleSubConnStateChange` 方法中，其方法调用链如下：

`ac.updateConnectivityState` -> `grpc.ClientConn.handleSubConnStateChange` -> `ccBalancerWrapper.handleSubConnStateChange`

来看一下最后一个方法 `ccBalancerWrapper.handleSubConnStateChange` 的源码：
```go
func (ccb *ccBalancerWrapper) handleSubConnStateChange(sc balancer.SubConn, s connectivity.State, err error) {
	if sc == nil {
		return
	}
	ccb.updateCh.Put(&scStateUpdate{
		sc:    sc,
		state: s,
		err:   err,
	})
}
```

该方法把一个 `balancer.SubConn` 和 `connectivity.State` 丢进了一个切片，然后通过一个 channel 控制另一个 goroutine 取数据
```go
func (b *Unbounded) Put(t interface{}) {
	b.mu.Lock()
	if len(b.backlog) == 0 {
		select {
		case b.c <- t:
			b.mu.Unlock()
			return
		default:
		}
	}
	b.backlog = append(b.backlog, t)
	b.mu.Unlock()
}
```

这里的数据写入后，在哪里读取，这就回到上文需要大家重点记住的一个 goroutine 调用了，还记得吗，试着回忆一下，没错就是 `go ccb.watcher()`

我们来看看 `watcher` 方法，由上文可知，我们写如的数据是 `scStateUpdate` 对象，因此如下源码就仅看获取该对象的 case 即可，省略了暂时不需要关注的代码：
```go
func (ccb *ccBalancerWrapper) watcher() {
	for {
		select {
		case t := <-ccb.updateCh.Get():
			ccb.updateCh.Load()
			if ccb.closed.HasFired() {
				break
			}
			switch u := t.(type) {
			case *scStateUpdate:
				ccb.balancerMu.Lock()
				ccb.balancer.UpdateSubConnState(u.sc, balancer.SubConnState{ConnectivityState: u.state, ConnectionError: u.err})
				ccb.balancerMu.Unlock()
			case ...:
				...
			default:
				logger.Errorf("ccBalancerWrapper.watcher: unknown update %+v, type %T", t, t)
			}
		case <-ccb.closed.Done():
		}
        ...
	}
}

```
由源码得知，其最终调用了 `balancer.Balancer.UpdateSubConnState` 方法，我们以 `roundrobin` 负载均衡器来查看，由上文知，gRPC 的 `balancer` 最终实现类是
`baseBalancer`， 因此 `balancer.Balancer.UpdateSubConnState` 最终落到了 `baseBalancer.UpdateSubConnState` 方法上，
```go
func (b *baseBalancer) UpdateSubConnState(sc balancer.SubConn, state balancer.SubConnState) {
    s := state.ConnectivityState
    ...
    oldS, ok := b.scStates[sc]
    if !ok {
        ...
        return
    }
    if oldS == connectivity.TransientFailure &&
        (s == connectivity.Connecting || s == connectivity.Idle) {
        if s == connectivity.Idle {
            sc.Connect()
        }
        return
    }
    b.scStates[sc] = s
    switch s {
    case connectivity.Idle:
        sc.Connect()
    case connectivity.Shutdown:
        // When an address was removed by resolver, b called RemoveSubConn but
        // kept the sc's state in scStates. Remove state for this sc here.
        delete(b.scStates, sc)
    case connectivity.TransientFailure:
        // Save error to be reported via picker.
        b.connErr = state.ConnectionError
    }
    
    b.state = b.csEvltr.RecordTransition(oldS, s)
    ...
    if (s == connectivity.Ready) != (oldS == connectivity.Ready) ||
        b.state == connectivity.TransientFailure {
        b.regeneratePicker()
    }
    b.cc.UpdateState(balancer.State{ConnectivityState: b.state, Picker: b.picker})
}
```
该方法中最终只会有状态 `connectivity.Ready` 的 `SubConn` 往下走，其他的状态要么被重新发起 `Connect`，要么被移出
最后一行代码发起 `balancer.ClientConn.UpdateState` 调用，因为 `ccBalancerWrapper` 为 `balancer.ClientConn` 的实现，因此来到
`balancer.ClientConn.UpdateState` 下，该方法做了两件事情：
- 更新 `balancer.Picker`
- 调用 `grpc.connectivityStateManager.updateState` 方法，该方法释放一个 channel 信号，通知 goroutine 进行信息处理，该 goroutine 我们后续再讲。

上文讲了这么多，那么负载算法在哪里，又何时调用呢？
由上文可知，`baseBalancer.UpdateSubConnState` 更新了一个 `picker`，这个 `picker` 来自哪里？追溯一下源码结合 `roundrobin` 负载均衡器可知，该 `picker`是在
`balancer.Builder` 的实现类调用 `base.NewBalancerBuilder` 创建实例时传入的 `base.PickBuilder` 实现类 `rrPickerBuilder` 构造出来的，看一下 `rrPickerBuilder`
的源码可知 `Pick` 方法中就是对 `SubConn` 进行负载算法的具体逻辑了。
```go
func (p *rrPicker) Pick(balancer.PickInfo) (balancer.PickResult, error) {
    p.mu.Lock()
    sc := p.subConns[p.next]
    p.next = (p.next + 1) % len(p.subConns)
    p.mu.Unlock()
    return balancer.PickResult{SubConn: sc}, nil
}
```

那么该方法什么时候调用呢？这里直接给出答案，在 `grpc.ClientConn` 发起 `Invoke` 方法调用时会通过调用链调用到，我们后续源码阅读到那里在来分析。

## 自定义负载均衡器
自定义负载均衡器首先需要了解  gRPC 的插件式编程，这部分内容可以自行 google。

### 环境
etcd
go

### 负载均衡目标
随机选择

1. 实现 `balancer.Builder`
我们就不一一实现其方法了，因为负载均衡器的重点在负载均衡算法，即实现 `base.PickerBuilder`，我们直接用 gRPC 提供的 `base.NewBalancerBuilder` 来创建 `balancer.Builder`
```go
const Name = "random"

func init() {
    balancer.Register(newBuilder())
}

func newBuilder() balancer.Builder {
    return base.NewBalancerBuilder(Name, &randomPickerBuilder{}, base.Config{HealthCheck: true})
}
```

2. 实现 `base.PickerBuilder`
```go
func (r *randomPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
    if len(info.ReadySCs) == 0 {
        return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
    }
    readyScs := make([]Conn, 0, len(info.ReadySCs))
    for sc, info := range info.ReadySCs {
        readyScs = append(readyScs, Conn{
            SubConn:     sc,
            SubConnInfo: info,
        })
    }
    return &randomPicker{
        subConns: readyScs,
        r:        rand.New(rand.NewSource(time.Now().UnixNano())),
    }
}
```

3. 实现 `balancer.Picker`
`balancer.Picker` 才是我们需要扩展的逻辑，即按照自己想要的负载均衡算法从 `SunConn` 列表中选择一个可用的 `SubConn` 创建链接。

```go
func (r *randomPicker) Pick(_ balancer.PickInfo) (balancer.PickResult, error) {
    next := r.r.Int() % len(r.subConns)
    sc := r.subConns[next]
    fmt.Printf("picked: %+v\n", sc.SubConnInfo.Address.Addr)
    return balancer.PickResult{
        SubConn: sc.SubConn,
    }, nil
}
```

4. 使用自定义负载均衡器
```go
r := resolverBuilder.NewCustomBuilder(resolverBuilder.Scheme)
options := []grpc.DialOption{grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(builder.Name)}
conn, err := grpc.Dial(resolverBuilder.Format("grpc-server"), options...)
```

### 演示效果
1. 启动多个 server 实例，我这里启动了三个
```shell
$ go run server.go -addr localhost:8888
```
```shell
$ go run server.go -addr localhost:8889
```
```shell
$ go run server.go -addr localhost:8890
```

2. 多次启动 client，观察 Pick 的日志输出
```shell
go run client.go
endpoints:  [localhost:8888 localhost:8889 localhost:8888 localhost:8889 localhost:8890]
picked: localhost:8888
output:  hi
```

```shell
go run client.go
endpoints:  [localhost:8888 localhost:8889 localhost:8888 localhost:8889 localhost:8890]
picked: localhost:8890
output:  hi
```

```shell
go run client.go
endpoints:  [localhost:8888 localhost:8889 localhost:8888 localhost:8889 localhost:8890]
picked: localhost:8889
output:  hi
```
...

## 总结
grpc 通过服务发现或者直连形式获取到 gRPC server 的实例的 endpoints，然后通知负载均衡器进行 `SubConn` 更新，对于新加入的 endpoint 进行实例创建，移出废弃的 endpoint，
最后通过状态更新将状态为 `Idle` 的 `SubConn` 进行管理，gRPC 在调用 `Connect`时，则会通过负载均衡器中的 `Picker` 去按照某一个负载均衡算法选择一个 `SubConn`
创建链接，如果创建成功则不再进行其他 `SubConn` 的尝试，否则会按照一定的退避算法进行重试，直到退避失败或者创建链接成功为止。

自定义负载均衡器的核心逻辑在于对 `Picker` 的实现，从 `SubConn` 列表中按照负载均衡算法选择一个 `SubConn` 创建链接，自定义负载均衡器和 `Resolver` 一样都是用到了插件式编程提供了扩展能力。

本次源码阅读只是为了理解 gRPC 的调用流程，其中有很多细节在源码注释中有说明，其可以加深我们对 gRPC 的理解，因此在理解本文介绍后，可以再次阅读源码加深一下理解。

## 源码
https://github.com/anqiansong/golang-notes/tree/main/example/grpc/balancer




