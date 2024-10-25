package main

import (
	"context"
	"fmt"
	"github.com/ShowyQuasar88/rpc_demo/go_grpc/pb/person"
	"google.golang.org/grpc"
)

func main() {
	// 客户端三步走
	// 1. 创建和服务器的连接
	// 2. 获取想要的服务
	// 3. 调用服务中的方法，获取返回
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
