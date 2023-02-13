package znet

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/alfmunny/willego/utils"
	"github.com/alfmunny/willego/zinterface"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	// DataLen uint32 (4 byte) + ID uint32(4 byte)
	return 8

}
func (d *DataPack) Pack(msg zinterface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataID()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func (d *DataPack) Unpack(binaryData []byte) (zinterface.IMessage, error) {
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	if utils.Config.MaxPackageSize > 0 && msg.DataLen > utils.Config.MaxPackageSize {
		return nil, errors.New("too large msg data recv!")
	}

	return msg, nil
}
