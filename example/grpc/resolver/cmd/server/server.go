package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"github.com/anqiansong/golang-notes/grpc/echo"
	"github.com/anqiansong/golang-notes/grpc/pkg/errorx"
	"github.com/anqiansong/golang-notes/grpc/pool"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":1001", "")

func main() {
	flag.Parse()
	pool.Register("grpc-server", *addr)
	lis, err := net.Listen("tcp", *addr)
	errorx.MustFatalln(err)
	s := grpc.NewServer()
	echo.RegisterEchoServiceServer(s, &server{})
	fmt.Println("serve ", *addr)
	err = s.Serve(lis)
	errorx.MustFatalln(err)
}

type server struct {
	echo.UnimplementedEchoServiceServer
}

func (s *server) Echo(ctx context.Context, req *echo.EchoIn) (*echo.EchoOut, error) {
	return &echo.EchoOut{
		Msg: req.In,
	}, nil
}
