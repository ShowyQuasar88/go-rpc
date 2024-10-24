package main

import (
	"fmt"
	"net/rpc"
)

func main() {
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
