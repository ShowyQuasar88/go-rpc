package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello, " + request
	fmt.Println(*reply)
	return nil
}

func main() {
	// 实例化 server
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     w,
			ReadCloser: r.Body,
		}
		fmt.Println(conn)
		_ = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	// 注册业务逻辑 Handler
	_ = rpc.RegisterName("HelloServiceRPC", &HelloService{})
	// 启动服务
	for {
		err := http.ListenAndServe(":1234", nil)
		if err != nil {
			panic(err)
		}
	}
}
