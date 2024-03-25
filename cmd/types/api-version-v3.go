package types

import (
	"bytes"
	"encoding/binary"
)

func (kafka *Kafka) APIVersion() *bytes.Buffer {

	request := bytes.NewBuffer([]byte{})
	request.Write(Uint32bytes(0))  // Request Size
	request.Write(Uint16bytes(18)) // API KEY
	request.Write(Uint16bytes(3))  // API Version
	request.Write(Uint32bytes(1))  // Correlation ID

	request.Write(Uint16bytes(len(kafka.ClientId))) // Client
	request.WriteString(kafka.ClientId)
	request.Write([]byte{0})

	request.Write([]byte{uint8(len(kafka.Name) + 1)})
	request.WriteString(kafka.Name)

	request.Write([]byte{uint8(len(kafka.Version) + 1)})
	request.WriteString(kafka.Version)

	request.Write([]byte{0})

	binary.BigEndian.PutUint32(request.Bytes()[0:], uint32(request.Len()-4))
	return request
}
