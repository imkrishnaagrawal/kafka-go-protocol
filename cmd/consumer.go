package main

import (
	"bytes"
	"encoding/binary"
)

const (
	REQUEST_PRODUCE      = 0
	REQUEST_FETCH        = 1
	REQUEST_MULTIFETCH   = 2
	REQUEST_MULTIPRODUCE = 3
	REQUEST_OFFSETS      = 4
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

func (kafka *Kafka) Fetch(topic string) (*FetchResponse, error) {

	request := bytes.NewBuffer([]byte{})
	request.Write(Uint32bytes(0))                   // Request Size
	request.Write(Uint16bytes(1))                   // API KEY
	request.Write(Uint16bytes(11))                  // API Version
	request.Write(Uint32bytes(7))                   // Correlation ID
	request.Write(Uint16bytes(len(kafka.ClientId))) // Client
	request.WriteString(kafka.ClientId)
	request.Write(Uint32bytes(int(kafka.ReplicaId)))
	request.Write(Uint32bytes(500))
	request.Write(Uint32bytes(1))
	request.Write(Uint32bytes(52428800))
	request.Write([]byte{1})
	request.Write(Uint32bytes(0))
	request.Write(Uint32bytes(-1))

	request.Write(Uint32bytes(1))

	request.Write(Uint16bytes(len(topic)))
	request.WriteString(topic)

	request.Write(Uint32bytes(1)) // parti len

	request.Write(Uint32bytes(0)) // partId
	request.Write(Uint32bytes(0)) // Leader Epoch

	request.Write(Uint32bytes(0))
	request.Write(Uint32bytes(0))

	request.Write(Uint32bytes(-1))
	request.Write(Uint32bytes(-1))

	request.Write(Uint32bytes(1048576))
	request.Write(Uint32bytes(0))
	request.Write(Uint16bytes(0))

	fetchRequest := kafka.NewFetchRequestV11(topic)
	err := kafka.Write(fetchRequest.Byte())
	if err != nil {
		return nil, err
	}
	reader, err := kafka.Read()
	if err != nil {
		return nil, err
	}

	fetchResponse := FetchResponse{}
	fetchResponse.Read(reader, kafka)
	return &fetchResponse, nil
}
