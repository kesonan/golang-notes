package builder

import (
	"fmt"

	"github.com/anqiansong/golang-notes/grpc/pool"
	"google.golang.org/grpc/resolver"
)

const Scheme = "custom"

func init() {
	resolver.Register(&customBuilder{})
}

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
	fmt.Println("endpoints: ", hosts)
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

func Format(addr string) string {
	return fmt.Sprintf("%s://%s", Scheme, addr)
}

type nopResolver struct {
}

func (*nopResolver) ResolveNow(resolver.ResolveNowOptions) {}

func (*nopResolver) Close() {}
