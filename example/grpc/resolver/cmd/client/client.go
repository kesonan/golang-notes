package main

import (
	"context"
	"fmt"

	"github.com/anqiansong/golang-notes/grpc/echo"
	"github.com/anqiansong/golang-notes/grpc/pkg/errorx"
	"github.com/anqiansong/golang-notes/grpc/resolver/builder"
	"google.golang.org/grpc"
)

func main() {
	r := builder.NewCustomBuilder(builder.Scheme)
	conn, err := grpc.Dial(builder.Format("grpc-server"), grpc.WithInsecure(), grpc.WithResolvers(r))
	errorx.MustFatalln(err)
	defer conn.Close()
	client := echo.NewEchoServiceClient(conn)
	ctx := context.Background()
	out, err := client.Echo(ctx, &echo.EchoIn{In: "hi"})
	errorx.MustFatalln(err)
	fmt.Println("output: ", out.Msg)
}
