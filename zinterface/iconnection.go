package zinterface

import "net"

// IConnection ...
type IConnection interface {
	// Start
	Start()

	//Stop
	Stop()

	// Get socket conn
	GetTCPConnection() *net.TCPConn

	GetConnID() uint32

	RemoteAddr() net.Addr

	Send(data []byte) error
}

// HandleFunc ...
type HandleFunc func(*net.TCPConn, []byte, int) error
