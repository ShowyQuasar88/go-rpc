package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func useJsonCodec() {
	conn, _ := net.Dial("tcp", "127.0.0.1:1234")
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	_ = client.Call("HelloServiceRPC.Hello", "showyquasar", &reply)
	fmt.Println(reply)
}

func originCall() {
	// 建立连接
	client, _ := rpc.Dial("tcp", "localhost:1234")
	// 选择要调用的服务名称
	var reply string
	err := client.Call("HelloServiceRPC.Hello", "bobby", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}

func main() {
	useJsonCodec()
}
