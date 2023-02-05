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
	// Router
	Router zinterface.IRouter
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

	var cid uint32
	cid = 0

	fmt.Println("start server succ, ", s.Name, " listening")
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		dealConn := NewConnection(conn, cid, s.Router)
		cid++

		go dealConn.Start()
	}
}

// Stop ...
func (s *Server) Stop() {

}

func (s *Server) AddRouter(router zinterface.IRouter) {
	s.Router = router
	fmt.Println("Add router succ!")
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
		Router:    nil,
	}
	return s
}
