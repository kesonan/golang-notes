# Resolver
gRPC 插件式编程之Resolver

--- 

随着微服务越来越盛行，服务间的通信也是绕不开的话题，gRPC 在众多 RPC 框架中算得上佼佼者，不仅其有一个好爸爸，grpc 在扩展方面也给开发者留
有足够的空间，今天我们将走进grpc 扩展之 Resolver，gRPC Resolver 提供了用户自行解析主机的扩展能力，我们在使用 gRPC 时，大家有没有想过，
为什么 gRPC 为什么支持以下几种格式的 target：

- 直连， 链接 target 为目标服务的endpoint
- dns 服务发现
- unix

其中在进入连接之前，gRPC 会根据用户是否提供了 Resolver 来进行目标服务的 endpoint 解析，今天我们来尝试写一个最简单的 etcd 做服务发现的例子

## 说明
源码阅读的 gRPC 版本为 `3.5.1`

## 环境
- etcd 安装
- go

## 思路
- 我们将为 server 服务，假设名称为 `grpc-server` 启动多个实例
- 以 `grpc-server` 为 key 向 etcd put 每个实例的 endpoint
  - 真正进入 etcd 的 key 为以 `grpc-server` + `/` + `随机值`
- 实现 resolver.Builder， 获取 `target`
- 从 etcd 读取以 `grpc-server` 为 prefix 的 endpoints
- 通知负载均衡器重新 pick 实例

## 实现
实现 resolver.Builder

```go
type customBuilder struct {
    scheme string
}

func NewCustomBuilder(scheme string) resolver.Builder {
  return &customBuilder{
    scheme: scheme,
  }
}

func (b *customBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
  var address []resolver.Address
  key := target.URL.Host
  hosts := pool.GetOr(key, nil)
  fmt.Println(hosts)
  for _, host := range hosts {
    address = append(address, resolver.Address{
        Addr: host,
    })
  }
  cc.UpdateState(resolver.State{Addresses: address})
  return &nopResolver{}, nil
}

func (b *customBuilder) Scheme() string {
	return Scheme
}
```

## 应用
在 client 发起调用时通过 `grpc.WithResolvers` DialOption 告知 gRPC
```go
r := builder.NewCustomBuilder(builder.Scheme)
	conn, err := grpc.Dial(builder.Format("grpc-server"), grpc.WithInsecure(), grpc.WithResolvers(r))
```

grpc.Dial 的第一个参数为 `target`，因此 `target` 并非一定是目标服务的 endpoint（仅直连模式才传目标服务的真正 endpoint），也可能是
指向某一个注册中心的遵循 URL 地址规范的一个值，便于开发者自定义 resolver.Builder 根据 `target` 拿到相应信息去做目标服务真正的 endpoints 解析，
如上文的 resolver.Builder 实现方法 
```go
Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error)
```
时，则根据解析后的 `resolver.Target` 拿到 etcd 的 key，再去获取目标服务的 endpoints，至于解析完后怎么通知负载均衡器的我们后续再讲。

## 示例结果
- 启动 etcd
```go
$ etcd
```

- 启动两个 server
```go
go run server.go -addr localhost:8888
```
```go
go run server.go -addr localhost:8889
```

- 启动 client
```go
$ go run client.go
endpoints: [localhost:8888 localhost:8889]
output: hi
```

## 原理
我们从 client 通过 `grpc.WithResolvers` 告知 gRPC resolver.Builder 后，他是怎么调用我们给的 resolver 的？
顺着 `grpc.DialContext` 源码看下去就知道，gRPC 会调用一个 `ClientConn.parseTargetAndFindResolver` 的方法，该
方法做了两个工作：

- 将 `grpc.DialContext` 的 `target` string值通过 `parseTarget` 解析为 `resolver.Target`，新版本为 `URL` 的包装体
- 寻找 resolver
  - gRPC 会优先从 `resolver.Target` 中的获取 `scheme` 名称，该值即为开发者在实现 `resolver.Builder` 时 `Scheme() string` 方法返回值一样。
  - gRPC 去 DialOption 中的 resolver 列表寻找名称相同 resolver
  - 通过 `newCCResolverWrapper` 方法调用 `resolver.Buidler.Build(target Target, cc ClientConn, opts BuildOptions) (Resolver, error)`方法实现解析
  - 告知负载均衡器后续处理

### 分解
1. gRPC 怎么知道该获取那个 resolver.Builder？
我们在通过 `grpc.DialContext` 传递的 `target` 一定要符合 [gRPC规范](https://github.com/grpc/grpc/blob/master/doc/naming.md)
其实就是符合 `URL` 格式，如 `http://foo.com`，`http` 即为 scheme，我们这里 `target` 格式为 `custom://xxx`，`xxx` 为 server 端的实例名称（即 `grpc-server`）
这样我们就告诉了 gRPC 该选择 `scheme` 为 `custom` 的 resolver.Builder。

2. gRPC 在哪里找到 resolver.Builder 的？
gRPC 的 resolver.Builder 并不会无中生有，而是我们在实例化 `resolver.Buidler` 的实现类时进行注册了，其实就是写到 gRPC 内部的一个全局 map 变量中了，
gRPC 在寻找是也是通 `scheme` 为 key 去这个 map 里找。

- 注册 resolver.Builder
```go
func init() {
    resolver.Register(&customBuilder{})
}
```
```go
func Register(b Builder) {
    m[b.Scheme()] = b
}
```

```go
var (
    // m is a map from scheme to resolver builder.
    m = make(map[string]Builder)
    // defaultScheme is the default scheme to use.
    defaultScheme = "passthrough"
)

```
- 获取 resolver.Builder

```go
for _, rb := range cc.dopts.resolvers {
    if scheme == rb.Scheme() {
        return rb
    }
}
return resolver.Get(scheme)
```

3. 解析 `target` 为 `resolver.Target`
```go
func parseTarget(target string) (resolver.Target, error) {
  u, err := url.Parse(target)
  if err != nil {
      return resolver.Target{}, err
  }
  
  endpoint := u.Path
  if endpoint == "" {
      endpoint = u.Opaque
  }
  endpoint = strings.TrimPrefix(endpoint, "/")
  return resolver.Target{
      Scheme:    u.Scheme,
      Authority: u.Host,
      Endpoint:  endpoint,
      URL:       *u,
  }, nil
}
```

## 源码位置
`clientconn.go:1622`

## 其他
在 gRPC 中，像这种 register & get 的形式成为插件式编程，通过这种手段给开发者提供了扩展入口，除 `resolver` 外，
`balancer`、 `compressor` 等都利用了这个手段提供了扩展入口，后续我们再来讨论。

## 总结
gRPC 的 `grpc.Dial` 或者 `gprc.DialContext` 的 `target` 并非一定是 server 的 endpoint，也有可能是满足开发者需求的
某类符合 `URL` 命名风格的值，如本 demo 中是 server 服务向 etcd 数据库存储 server 启动的每个实例的 endpoint 的 `key` 的 `prefix`，
如果用户没有提供 `resolver.Builder` ， gRPC 会根据默认 `scheme` （`resolver.GetDefaultScheme()`） 去查找。

## 源码
https://github.com/anqiansong/golang-notes/tree/main/example/grpc/resolver