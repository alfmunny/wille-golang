package znet

import "github.com/alfmunny/willego/zinterface"

type Request struct {
	conn zinterface.IConnection
	data []byte
}

func (r *Request) GetConnection() zinterface.IConnection {
	return r.conn

}

func (r *Request) GetData() []byte {
	return r.data
}
