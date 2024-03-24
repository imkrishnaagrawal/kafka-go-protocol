package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

type Consumer struct {
	ClientId  string
	ReplicaId int
}

type Producer struct {
}

type Kafka struct {
	conn *net.TCPConn

	Name    string
	Version string
	*Consumer
	*Producer
	*ResponseErrorSet
}

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

type Topic struct {
	NameLength      uint16
	Name            string
	PartitionLength uint32
	Partitions      []Partition
}

func (topic *Topic) Read(reader *bytes.Reader, conn *net.TCPConn) {
	binary.Read(reader, binary.BigEndian, &topic.NameLength)
	topic.Name = ReadString(reader, int(topic.NameLength))
	binary.Read(reader, binary.BigEndian, &topic.PartitionLength)
	for ii := 0; ii < int(topic.PartitionLength); ii++ {
		partition := Partition{}
		partition.Read(reader, conn)
		topic.Partitions = append(topic.Partitions, partition)
	}
}

type Partition struct {
	Id                        uint32
	ErrorCode                 uint16
	Offset                    uint64
	LastStableOffset          uint64
	LogStartOffset            uint64
	AbortedTransactionsLength uint32
	AbortedTransactions       []AbortedTransaction
	ReplicaID                 int32
	MessageSetSize            uint32
	MessageSet                MessageSet
}

func (partition *Partition) Read(reader *bytes.Reader, conn *net.TCPConn) {
	binary.Read(reader, binary.BigEndian, &partition.Id)
	binary.Read(reader, binary.BigEndian, &partition.ErrorCode)
	binary.Read(reader, binary.BigEndian, &partition.Offset)
	binary.Read(reader, binary.BigEndian, &partition.LastStableOffset)
	binary.Read(reader, binary.BigEndian, &partition.LogStartOffset)
	binary.Read(reader, binary.BigEndian, &partition.AbortedTransactionsLength)
	for iii := 0; iii < int(partition.AbortedTransactionsLength); iii++ {
		abortedTransaction := AbortedTransaction{}
		abortedTransaction.Read(reader)
		partition.AbortedTransactions = append(partition.AbortedTransactions, abortedTransaction)
	}
	binary.Read(reader, binary.BigEndian, &partition.ReplicaID)
	binary.Read(reader, binary.BigEndian, &partition.MessageSetSize)

	messageBuf := make([]byte, partition.MessageSetSize)
	conn.Read(messageBuf)
	messageReader := bytes.NewReader(messageBuf)

	messageSet := MessageSet{}
	messageSet.Read(messageReader)
	partition.MessageSet = messageSet

}

type AbortedTransaction struct {
	ProducerId  uint64
	FirstOffset uint64
}

func (abortedTransaction *AbortedTransaction) Read(reader *bytes.Reader) {
	binary.Read(reader, binary.BigEndian, &abortedTransaction.ProducerId)
	binary.Read(reader, binary.BigEndian, &abortedTransaction.FirstOffset)
}

type MessageSet struct {
	RecordBatch
}

func (messageSet *MessageSet) Read(messageReader *bytes.Reader) {
	messageSet.RecordBatch = RecordBatch{}
	messageSet.RecordBatch.Read(messageReader)
	binary.Read(messageReader, binary.BigEndian, &messageSet.RecordBatch.RecordLength)

	record := Record{}
	record.Read(messageReader)
	binary.Read(messageReader, binary.BigEndian, record.HeaderLength)
	for h := 0; h < int(record.HeaderLength); h++ {
		header := Header{}
		header.Read(messageReader)
		record.Headers = append(record.Headers, header)
	}
	messageSet.RecordBatch.Records = append(messageSet.RecordBatch.Records, record)
}

type RecordBatch struct {
	BaseOffset           int64
	BatchLength          int32
	PartitionLeaderEpoch int32
	MagicByte            int8
	CRC                  uint32
	Attributes           int16
	// CompressionCodec 0-2
	// TimestampType    3
	// isTransactional  4
	// isControlBatch   5
	// hasDeleteHorizonMs 6
	// Unused 7-15
	LastOffsetDelta int32
	BaseTimestamp   int64
	MaxTimestamp    int64
	ProducerID      int64
	ProducerEpoch   int16
	BaseSequence    int32
	RecordLength    int32
	Records         []Record
}

func (recordBatch *RecordBatch) Read(reader *bytes.Reader) {
	binary.Read(reader, binary.BigEndian, &recordBatch.BaseOffset)
	binary.Read(reader, binary.BigEndian, &recordBatch.BatchLength)
	binary.Read(reader, binary.BigEndian, &recordBatch.PartitionLeaderEpoch)
	binary.Read(reader, binary.BigEndian, &recordBatch.MagicByte)
	binary.Read(reader, binary.BigEndian, &recordBatch.CRC)
	binary.Read(reader, binary.BigEndian, &recordBatch.Attributes)
	binary.Read(reader, binary.BigEndian, &recordBatch.LastOffsetDelta)
	binary.Read(reader, binary.BigEndian, &recordBatch.BaseTimestamp)
	binary.Read(reader, binary.BigEndian, &recordBatch.MaxTimestamp)
	binary.Read(reader, binary.BigEndian, &recordBatch.ProducerID)
	binary.Read(reader, binary.BigEndian, &recordBatch.ProducerEpoch)
	binary.Read(reader, binary.BigEndian, &recordBatch.BaseSequence)
}

type Record struct {
	RecordSize     int64
	Attributes     int8
	TimestampDelta int64
	OffsetDelta    int64
	KeyLength      int64
	Key            string
	ValueLength    int64
	Value          string
	HeaderLength   uint32
	Headers        []Header
}

func (record *Record) Read(messageReader *bytes.Reader) error {
	var err error
	record.RecordSize, err = binary.ReadVarint(messageReader)
	if err != nil {
		return fmt.Errorf("error reading record size varint")
	}
	binary.Read(messageReader, binary.BigEndian, &record.Attributes)
	record.TimestampDelta, err = binary.ReadVarint(messageReader)
	if err != nil {
		return fmt.Errorf("error reading timestamp delta varint")
	}
	record.OffsetDelta, err = binary.ReadVarint(messageReader)
	if err != nil {
		return fmt.Errorf("error offset delta varint")
	}
	record.KeyLength, err = binary.ReadVarint(messageReader)
	if err != nil {
		return fmt.Errorf("error key length varint")
	}
	record.Key = ReadString(messageReader, int(record.KeyLength))

	record.ValueLength, err = binary.ReadVarint(messageReader)
	if err != nil {
		return fmt.Errorf("error reading record value length  varint")
	}
	record.Value = ReadString(messageReader, int(record.ValueLength))

	return nil
}

type Header struct {
	HeaderKeyLength   int64
	HeaderKey         string
	HeaderValueLength int64
	HeaderValue       string
}

func (header *Header) Read(reader *bytes.Reader) error {
	var err error
	header.HeaderKeyLength, err = binary.ReadVarint(reader)
	if err != nil {
		return fmt.Errorf("error reading header key length varint")
	}
	header.HeaderKey = ReadString(reader, int(header.HeaderKeyLength))
	header.HeaderValueLength, err = binary.ReadVarint(reader)
	if err != nil {
		return fmt.Errorf("error reading header value length varint")
	}
	header.HeaderValue = ReadString(reader, int(header.HeaderValueLength))
	return nil
}

func ReadString(reader io.Reader, length int) string {
	name := make([]byte, length)
	io.ReadFull(reader, name)
	return string(name)
}

func PrintBytesInHexFormat(data []byte) {
	for _, b := range data {
		fmt.Printf("\\x%02x", b)
	}
	fmt.Println()
}
