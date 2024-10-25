package main

import (
	"context"
	"fmt"
	"github.com/ShowyQuasar88/rpc_demo/go_grpc/pb/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"time"
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
func (*PersonServer) SearchIn(server grpc.ClientStreamingServer[person.PersonReq, person.PersonResp]) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err == io.EOF {
			err := server.SendAndClose(&person.PersonResp{Name: "Done"})
			if err != nil {
				return err
			}
			break
		} else if err != nil {
			err := server.SendAndClose(&person.PersonResp{Name: fmt.Sprintf("err: %v", err)})
			if err != nil {
				return err
			}
			break
		}
	}
	//return status.Errorf(codes.Unimplemented, "method SearchIn not implemented")
	return nil
}
func (*PersonServer) SearchOut(req *person.PersonReq, server grpc.ServerStreamingServer[person.PersonResp]) error {
	name := req.Name
	for idx := 0; idx < 10; idx++ {
		err := server.Send(&person.PersonResp{Name: name, Age: int32(idx)})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
	//return status.Errorf(codes.Unimplemented, "method SearchOut not implemented")
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
