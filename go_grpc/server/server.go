package main

import (
	"context"
	"github.com/ShowyQuasar88/rpc_demo/go_grpc/pb/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type PersonServer struct {
	person.UnimplementedSearchServiceServer
}

func (*PersonServer) Search(ctx context.Context, req *person.PersonReq) (*person.PersonResp, error) {
	name := req.GetName()
	age := req.GetAge()
	resp := &person.PersonResp{
		Name: "receive: " + name,
		Age:  age,
	}
	//return resp, status.Errorf(codes.Unimplemented, "method Search not implemented")
	return resp, nil
}
func (*PersonServer) SearchIn(grpc.ClientStreamingServer[person.PersonReq, person.PersonResp]) error {
	return status.Errorf(codes.Unimplemented, "method SearchIn not implemented")
}
func (*PersonServer) SearchOut(*person.PersonReq, grpc.ServerStreamingServer[person.PersonResp]) error {
	return status.Errorf(codes.Unimplemented, "method SearchOut not implemented")
}
func (*PersonServer) SearchIO(grpc.BidiStreamingServer[person.PersonReq, person.PersonResp]) error {
	return status.Errorf(codes.Unimplemented, "method SearchIO not implemented")
}

func main() {
	// server 创建四部曲
	// 1. 选择监听的端口和协议
	// 2. 创建一个 grpc server
	// 3. 将服务注册到 server 中
	// 4. 启动并监听，进行服务
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &PersonServer{})
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
