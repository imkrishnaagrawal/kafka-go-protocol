package types

import (
	"bytes"
	"encoding/binary"
)

type FetchResponse struct {
	CorrelationID  uint32
	ThrottleTime   uint32
	ErrorCode      uint16
	FetchSessionID uint32
	TopicLength    uint32
	Topics         []Topic
}

func (fetchResponse *FetchResponse) Read(reader *bytes.Reader, kafka *Kafka) {

	binary.Read(reader, binary.BigEndian, &fetchResponse.CorrelationID)
	binary.Read(reader, binary.BigEndian, &fetchResponse.ThrottleTime)
	binary.Read(reader, binary.BigEndian, &fetchResponse.ErrorCode)
	binary.Read(reader, binary.BigEndian, &fetchResponse.FetchSessionID)
	binary.Read(reader, binary.BigEndian, &fetchResponse.TopicLength)
	for i := 0; i < int(fetchResponse.TopicLength); i++ {
		topic := Topic{}
		topic.Read(reader, kafka.conn)
		fetchResponse.Topics = append(fetchResponse.Topics, topic)
	}

}
