package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	hello_grpc "rpc_demo/protobuf_hello_world/proto"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) SayHi(ctx context.Context, req *hello_grpc.Req) (res *hello_grpc.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "服务端收到了消息并告诉你：Hello " + req.GetMessage()}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	err = s.Serve(l)
	if err != nil {
		log.Fatalf("s.Serve err: %v", err)
	}
}
