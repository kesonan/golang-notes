package builder

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
)

type randomPickerBuilder struct {
}

type Conn struct {
	SubConn     balancer.SubConn
	SubConnInfo base.SubConnInfo
}

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

type randomPicker struct {
	subConns []Conn
	mu       sync.Mutex
	r        *rand.Rand
}

func (r *randomPicker) Pick(_ balancer.PickInfo) (balancer.PickResult, error) {
	next := r.r.Int() % len(r.subConns)
	sc := r.subConns[next]
	fmt.Printf("picked: %+v\n", sc.SubConnInfo.Address.Addr)
	return balancer.PickResult{
		SubConn: sc.SubConn,
	}, nil
}
