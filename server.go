package main

import "github.com/alfmunny/willego/znet"

func main() {
	s := znet.NewServer("Server 1.0")
	s.Serve()
}
