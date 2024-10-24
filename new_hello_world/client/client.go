package main

import (
	"fmt"
	"rpc_demo/new_hello_world/client/proxy"
)

func originCall() {
	client, _ := proxy.NewHelloServiceClient("tcp", "localhost:1234")
	var reply string
	err := client.Hello("bobby", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}

func main() {
	originCall()
}
