package proxy

import (
	"net/rpc"
	"rpc_demo/new_hello_world/handler"
)

type IHelloService interface {
	Hello(request string, reply *string) error
}

func RegistryHelloService(srv IHelloService) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
