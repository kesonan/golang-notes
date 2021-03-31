package main

import (
	"context"
	greet "github.com/anqiansong/golang-notes/rpc"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8888")
	if err != nil {
		log.Fatal(err)
	}

	client := greet.NewGreeterClient(conn)
	// ignore returns
	_, _ = client.SayHello(context.Background(), &greet.HelloReq{})
}
