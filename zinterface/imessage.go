package zinterface

// IMessage is to include attribute of a message
type IMessage interface {
	GetDataID() uint32
	GetDataLen() uint32
	GetData() []byte

	SetDataID(uint32)
	SetDataLen(uint32)
	SetData([]byte)
}
