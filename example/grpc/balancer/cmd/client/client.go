package main

import (
	"context"
	"fmt"

	"github.com/anqiansong/golang-notes/grpc/balancer/builder"
	"github.com/anqiansong/golang-notes/grpc/echo"
	"github.com/anqiansong/golang-notes/grpc/pkg/errorx"
	resolverBuilder "github.com/anqiansong/golang-notes/grpc/resolver/builder"
	"google.golang.org/grpc"
)

func main() {
	r := resolverBuilder.NewCustomBuilder(resolverBuilder.Scheme)
	options := []grpc.DialOption{grpc.WithInsecure(), grpc.WithResolvers(r), grpc.WithBalancerName(builder.Name)}
	conn, err := grpc.Dial(resolverBuilder.Format("grpc-server"), options...)
	errorx.MustFatalln(err)
	defer conn.Close()
	client := echo.NewEchoServiceClient(conn)
	ctx := context.Background()
	out, err := client.Echo(ctx, &echo.EchoIn{In: "hi"})
	errorx.MustFatalln(err)
	fmt.Println("output: ", out.Msg)
}
