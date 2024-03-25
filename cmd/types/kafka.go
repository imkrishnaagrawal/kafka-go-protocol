package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type Consumer struct {
	ClientId  string
	ReplicaId int32
}

type Producer struct{}

type Kafka struct {
	conn    *net.TCPConn
	Name    string
	Version string
	*Consumer
	*Producer
	*ResponseErrorSet
}

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

func (kafka *Kafka) Write(request []byte) error {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, uint32(len(request)))
	binary.Write(buf, binary.BigEndian, request)
	_, err := kafka.conn.Write(buf.Bytes())
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
