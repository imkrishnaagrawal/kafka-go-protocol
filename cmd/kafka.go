package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func NewKafka() *Kafka {
	return &Kafka{
		Consumer: &Consumer{
			ClientId:  "custom",
			ReplicaId: -1,
		},
		Name:             "confluent-kafka-go",
		Version:          "2.3.0-custom-2.3.0",
		ResponseErrorSet: NewResponseErrorSet(),
	}
}

func (kafka *Kafka) Close() error {
	return kafka.conn.Close()
}

func (kafka *Kafka) Connect(host string) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		return fmt.Errorf("failed to resolve tcp address %s", err)
	}
	kafka.conn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to broker %s", err)
	}
	return nil
}

func (kafka *Kafka) Write(request *bytes.Buffer) error {
	binary.BigEndian.PutUint32(request.Bytes()[0:], uint32(request.Len()-4))

	_, err := kafka.conn.Write(request.Bytes())
	if err != nil {
		return fmt.Errorf("failed to send handshake: %s", err)
	}
	return nil
}

func (kafka *Kafka) Read() (*bytes.Reader, error) {
	lenBuf := make([]byte, 4)
	_, err := kafka.conn.Read(lenBuf)
	if err != nil {
		return nil, fmt.Errorf("failed to read response %s", err)
	}
	length := binary.BigEndian.Uint32(lenBuf)
	responseBuf := make([]byte, length)
	kafka.conn.Read(responseBuf)

	return bytes.NewReader(responseBuf), nil
}
