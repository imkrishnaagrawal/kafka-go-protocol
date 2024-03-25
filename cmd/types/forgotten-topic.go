package types

import (
	"bytes"
	"encoding/binary"
)

type ForgottenTopic struct {
	Topic      string
	Partitions int32
}

func (forgottenTopic *ForgottenTopic) Byte() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, uint16(len(forgottenTopic.Topic)))
	binary.Write(buf, binary.BigEndian, []byte(forgottenTopic.Topic))
	binary.Write(buf, binary.BigEndian, forgottenTopic.Partitions)
	return buf.Bytes()
}

type ForgottenTopicList []ForgottenTopic

func (topics ForgottenTopicList) Byte() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, uint32(len(topics)))
	for _, topic := range topics {
		binary.Write(buf, binary.BigEndian, topic.Byte())
	}
	return buf.Bytes()
}
