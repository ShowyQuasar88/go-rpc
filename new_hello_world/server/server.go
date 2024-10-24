package main

import (
	"net"
	"net/rpc"
	"rpc_demo/new_hello_world/handler"
	"rpc_demo/new_hello_world/server/proxy"
)

func main() {
	// 实例化 server
	listener, _ := net.Listen("tcp", ":1234")
	src := handler.HelloService{}
	// 注册业务逻辑 Handler
	_ = proxy.RegistryHelloService(&src)
	// 启动服务
	for {
		conn, _ := listener.Accept()
		rpc.ServeConn(conn)
	}
}
