package tcp

import (
	"bytes"
	"encoding/binary"
)

type DataPgk struct {
	Data []byte
	Len  uint32
}

func (d *DataPgk) Marshal() []byte {
	byteBuffer := bytes.NewBuffer([]byte{})
	binary.Write(byteBuffer, binary.BigEndian, d.Len)
	return append(byteBuffer.Bytes(), d.Data...)
}
