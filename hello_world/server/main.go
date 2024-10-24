package main

import (
	"fmt"
	"net"
	"net/rpc"
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
	conn, _ := listener.Accept()
	fmt.Println(conn)
	rpc.ServeConn(conn)
}
