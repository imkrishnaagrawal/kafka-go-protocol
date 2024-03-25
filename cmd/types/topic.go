package types

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Topic struct {
	TopicNameLength uint16
	TopicName       string
	PartitionLength uint32
	Partitions      []Partition
}

func (topic *Topic) Read(reader *bytes.Reader, conn *net.TCPConn) {
	binary.Read(reader, binary.BigEndian, &topic.TopicNameLength)
	topic.TopicName = ReadString(reader, int(topic.TopicNameLength))
	binary.Read(reader, binary.BigEndian, &topic.PartitionLength)
	for ii := 0; ii < int(topic.PartitionLength); ii++ {
		partition := Partition{}
		partition.Read(reader, conn)
		topic.Partitions = append(topic.Partitions, partition)
	}
}
