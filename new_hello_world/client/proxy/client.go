package proxy

import (
	"net/rpc"
	"rpc_demo/new_hello_world/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, address string) (*HelloServiceStub, error) {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceStub{conn}, nil
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
