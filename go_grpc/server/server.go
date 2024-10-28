package main

import (
	"context"
	"fmt"
	"github.com/ShowyQuasar88/rpc_demo/go_grpc/pb/person"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
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
func (*PersonServer) SearchIO(server grpc.BidiStreamingServer[person.PersonReq, person.PersonResp]) error {
	str := make(chan string) // 来一个为空的缓冲区
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			req, err := server.Recv()
			if err != nil {
				str <- "Done"
				break
			}
			str <- req.GetName()
		}
	}()
	go func() {
		defer wg.Done()
		for {
			receive := <-str
			err := server.Send(&person.PersonResp{Name: "Receive: " + receive})
			if err != nil {
				fmt.Println(err)
				break
			}
			if receive == "Done" {
				break
			}
		}
	}()
	wg.Wait()
	return nil
	//return status.Errorf(codes.Unimplemented, "method SearchIO not implemented")
}

func registerGateWay(wg *sync.WaitGroup) {
	defer wg.Done()

	mux := runtime.NewServeMux()

	gwServer := &http.Server{
		Handler: mux,
		Addr:    ":8090",
	}

	conn, err := grpc.NewClient("127.0.0.1:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	err = person.RegisterSearchServiceHandler(context.Background(), mux, conn)
	if err != nil {
		panic(err)
	}

	err = gwServer.ListenAndServe()
	if err != nil {
		return
	}
}

func registerGrpc(wg *sync.WaitGroup) {
	defer wg.Done()
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

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go registerGateWay(&wg)
	go registerGrpc(&wg)
	wg.Wait()
}
