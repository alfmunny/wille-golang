package znet

import (
	"fmt"
	"net"

	"github.com/alfmunny/willego/zinterface"
)

// Connection ...
type Connection struct {
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	ExitChan chan bool

	Router zinterface.IRouter
}

// NewConnection ...
func NewConnection(conn *net.TCPConn, connID uint32, router zinterface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

//Stop ...
func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnID)

	if c.isClosed == true {
		return
	}

	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) StartReader() {
	fmt.Println(" Reader Goroutine is running...")
	defer fmt.Println("connID = ", c.ConnID, " Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		req := Request{
			conn: c,
			data: buf,
		}

		go func(request zinterface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}

}

func (c *Connection) Start() {
	fmt.Println("Conn Start() ... ConnID = ", c.ConnID)
	go c.StartReader()

}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
