package znet

// Message is to include attribute of a message
type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

// GetDataID ...
func (m *Message) GetDataID() uint32 {
	return m.ID
}

// GetDataLen ...
func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

// GetData ...
func (m *Message) GetData() []byte {
	return m.Data
}

// SetDataID ...
func (m *Message) SetDataID(id uint32) {
	m.ID = id
}

// SetDataLen ...
func (m *Message) SetDataLen(dataLen uint32) {
	m.DataLen = dataLen
}

// SetData ...
func (m *Message) SetData(data []byte) {
	m.Data = data
}
