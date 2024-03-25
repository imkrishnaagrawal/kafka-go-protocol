package types

import (
	"bytes"
	"encoding/binary"
)

type RequestHeader struct {
	APIKey        int16
	APIVersion    int16
	CorrelationID int32
	ClientId      string
}

func (requestHeader *RequestHeader) Byte() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, requestHeader.APIKey)
	binary.Write(buf, binary.BigEndian, requestHeader.APIVersion)
	binary.Write(buf, binary.BigEndian, requestHeader.CorrelationID)
	binary.Write(buf, binary.BigEndian, uint16(len(requestHeader.ClientId)))
	binary.Write(buf, binary.BigEndian, []byte(requestHeader.ClientId))
	return buf.Bytes()
}
