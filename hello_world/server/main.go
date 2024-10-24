package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	return nil
}

func main() {
	// 实例化 server
	listener, _ := net.Listen("tcp", ":1234")
	// 注册业务逻辑 Handler
	_ = rpc.RegisterName("HelloServiceRPC", &HelloService{})
	// 启动服务
	for {
		conn, _ := listener.Accept()
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
