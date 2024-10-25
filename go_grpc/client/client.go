package main

import (
	"context"
	"fmt"
	"github.com/ShowyQuasar88/rpc_demo/go_grpc/pb/person"
	"google.golang.org/grpc"
	"io"
	"sync"
	"time"
)

// originSearch 调用服务端原始 Search 方法
func originSearch() {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := person.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &person.PersonReq{Name: "ShowyQuasar88", Age: 18})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

// searchIn 调用服务端 Search 流式传入方法
func searchIn() {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := person.NewSearchServiceClient(conn)
	stream, err := client.SearchIn(context.Background())
	if err != nil {
		panic(err)
	}
	for idx := 0; idx < 10; idx++ {
		err := stream.Send(&person.PersonReq{Name: "ShowyQuasar88", Age: int32(idx)})
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

// searchOut 调用服务端 Search 流式返回方法
func searchOut() {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := person.NewSearchServiceClient(conn)
	stream, err := client.SearchOut(context.Background(), &person.PersonReq{Name: "ShowyQuasar88", Age: 18})
	if err != nil {
		panic(err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(resp)
	}
}

// searchIO 调用服务端 Search 流式请求返回方法
func searchIO() {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := person.NewSearchServiceClient(conn)
	stream, err := client.SearchIO(context.Background())
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for idx := 0; idx < 10; idx++ {
			err := stream.Send(&person.PersonReq{Name: "ShowyQuasar88"})
			if err != nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			req, err := stream.Recv()
			if err != nil {
				break
			}
			fmt.Println(req)
		}
	}()
	wg.Wait()
}

func main() {
	// 客户端三步走
	// 1. 创建和服务器的连接
	// 2. 获取想要的服务
	// 3. 调用服务中的方法，获取返回
	//originSearch()
	//searchIn()
	//searchOut()
	searchIO()
}
