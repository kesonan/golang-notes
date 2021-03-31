package main

import (
	"context"
	greet "github.com/anqiansong/golang-notes/rpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GreetImplement struct {
	greet.UnimplementedGreeterServer
}

func (i *GreetImplement) SayHello(ctx context.Context, in *greet.HelloReq) (*greet.HelloReply, error) {
	return &greet.HelloReply{}, nil
}

func main() {
	server := grpc.NewServer()
	greet.RegisterGreeterServer(server, &GreetImplement{})
	lis, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	_ = server.Serve(lis)
}
