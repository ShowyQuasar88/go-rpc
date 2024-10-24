package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	hello_grpc "rpc_demo/protobuf_hello_world/proto"
)

func main() {
	conn, err := grpc.NewClient("localhost:8888", grpc.WithInsecure())
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	client := hello_grpc.NewHelloGRPCClient(conn)
	resp, err := client.SayHi(context.Background(), &hello_grpc.Req{Message: "客户端发送请求"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(resp.GetMessage())
}
