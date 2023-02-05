package main

import (
	"fmt"

	"github.com/alfmunny/willego/zinterface"
	"github.com/alfmunny/willego/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (router *PingRouter) PreHandle(r zinterface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := r.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
	if err != nil {
		fmt.Println("call back before ping error")
	}

}

func (router *PingRouter) Handle(r zinterface.IRequest) {
	fmt.Println("Call Router Handle")
	_, err := r.GetConnection().GetTCPConnection().Write([]byte("ping..."))
	if err != nil {
		fmt.Println("call back ping error")
	}

}

func (router *PingRouter) PostHandle(r zinterface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := r.GetConnection().GetTCPConnection().Write([]byte("afer ping..."))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}

func main() {
	s := znet.NewServer("Server 1.0")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
