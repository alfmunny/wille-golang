package znet

import (
	"fmt"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {

	listener, err := net.Listen("tcp", "127.0.0.1:7777")

	if err != nil {
		fmt.Println("server listen err: ", err)
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept error", err)
			}

			go func(conn net.Conn) {
				dp := NewDataPack()

				for {

				}

			}(conn)
		}

	}()

}
