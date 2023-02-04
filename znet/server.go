package znet

import (
	"fmt"
	"net"

	"github.com/alfmunny/willego/zinterface"
)

//Server ...
type Server struct {
	// name
	Name string
	// ip version
	IPVersion string
	// ip
	IP string
	// port
	Port int
}

// Start ...
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP :%s, Port %d, is starting\n", s.IP, s.Port)
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error : ", err)
	}

	listener, err := net.ListenTCP(s.IPVersion, addr)

	if err != nil {
		fmt.Println("listen ", s.IPVersion, " err ", err)
	}

	fmt.Println("start server succ, ", s.Name, " listening")
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		go func() {
			for {
				buf := make([]byte, 512)
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("recv buf err", err)
					continue
				}

				fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Println("write back buf err", err)
					continue
				}
			}

		}()
	}
}

// Stop ...
func (s *Server) Stop() {

}

// Serve ...
func (s *Server) Serve() {
	go s.Start()
	// Do something after server start
	select {}
}

// NewServer Create Server
func NewServer(name string) zinterface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      9998,
	}
	return s
}
