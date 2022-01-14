package pool

import (
	"context"
	"fmt"
	"time"

	"github.com/anqiansong/golang-notes/grpc/pkg/errorx"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type addrs = map[string]map[string]struct{}

var pools = make(addrs)

var client *clientv3.Client

func init() {
	var err error
	client, err = clientv3.NewFromURL("localhost:2379")
	errorx.MustFatalln(err)
}

func Register(key string, addr string) {
	key = fmt.Sprintf("%s/%v", key, time.Now().UnixNano())
	_, _ = client.Put(context.Background(), key, addr)
}

func GetOr(key string, dft []string) []string {
	resp, err := client.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return dft
	}
	var list []string
	for _, kv := range resp.Kvs {
		list = append(list, string(kv.Value))
	}
	return list
}
