package types

import (
	"bytes"
	"encoding/binary"
)

func (kafka *Kafka) Fetch(topic string) ([]FetchResponse, error) {

	fetchResponses := []FetchResponse{}
	currentOffset := 0
	correlationID := 7

	fetchRequest := kafka.NewFetchRequestV11(topic, int64(currentOffset), int32(correlationID))
	err := kafka.Write(fetchRequest.Byte())
	if err != nil {
		return nil, err
	}
	reader, err := kafka.Read()
	if err != nil {
		return nil, err
	}
	currentOffset += 2
	correlationID += 1
	fetchResponse := FetchResponse{}
	fetchResponse.Read(reader, kafka)

	// fmt.Printf("%+v\n", fetchResponse)
	fetchResponses = append(fetchResponses, fetchResponse)

	return fetchResponses, nil
}

type FetchRequestV11 struct {
	RequestHeader     RequestHeader
	ReplicaID         int32
	MaxWaitTime       int32
	MinBytes          int32
	MaxBytes          int32
	IsolationLevel    int8
	FetchSessionID    int32
	FetchSessionEpoch int32
	Topics            FetchTopicV11List
	ForgottenTopics   ForgottenTopicList
	RackId            string
}

func (kafka *Kafka) NewFetchRequestV11(topic string, offset int64, correlationID int32) *FetchRequestV11 {
	return &FetchRequestV11{
		RequestHeader: RequestHeader{
			APIKey:        API_KEY_FETCH,
			APIVersion:    11,
			CorrelationID: correlationID,
			ClientId:      kafka.ClientId,
		},
		ReplicaID:         kafka.ReplicaId,
		MaxWaitTime:       500,
		MinBytes:          1,
		MaxBytes:          52428800,
		IsolationLevel:    1,
		FetchSessionID:    0,
		FetchSessionEpoch: -1,
		Topics: FetchTopicV11List{
			FetchTopicV11{
				TopicName: topic,
				Partitions: FetchPartitionV11List{
					FetchPartitionV11{
						PartitionID:    0,
						LeaderEpoch:    0,
						Offset:         offset,
						LogStartOffset: -1,
						MaxBytes:       1048576,
					},
				},
			},
		},
	}
}

func (fetchRequest *FetchRequestV11) Byte() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, fetchRequest.RequestHeader.Byte())
	binary.Write(buf, binary.BigEndian, fetchRequest.ReplicaID)
	binary.Write(buf, binary.BigEndian, fetchRequest.MaxWaitTime)
	binary.Write(buf, binary.BigEndian, fetchRequest.MinBytes)
	binary.Write(buf, binary.BigEndian, fetchRequest.MaxBytes)
	binary.Write(buf, binary.BigEndian, fetchRequest.IsolationLevel)
	binary.Write(buf, binary.BigEndian, fetchRequest.FetchSessionID)
	binary.Write(buf, binary.BigEndian, fetchRequest.FetchSessionEpoch)
	binary.Write(buf, binary.BigEndian, fetchRequest.Topics.Byte())
	binary.Write(buf, binary.BigEndian, fetchRequest.ForgottenTopics.Byte())
	binary.Write(buf, binary.BigEndian, uint16(len(fetchRequest.RackId)))
	binary.Write(buf, binary.BigEndian, []byte(fetchRequest.RackId))
	return buf.Bytes()
}

type FetchTopicV11 struct {
	TopicName  string
	Partitions FetchPartitionV11List
}

type FetchTopicV11List []FetchTopicV11

func (fetchTopic *FetchTopicV11) Byte() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, uint16(len(fetchTopic.TopicName)))
	binary.Write(buf, binary.BigEndian, []byte(fetchTopic.TopicName))
	binary.Write(buf, binary.BigEndian, fetchTopic.Partitions.Byte())
	return buf.Bytes()
}

func (fetchTopics FetchTopicV11List) Byte() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, uint32(len(fetchTopics)))
	for _, topic := range fetchTopics {
		binary.Write(buf, binary.BigEndian, topic.Byte())
	}
	return buf.Bytes()
}

type FetchPartitionV11 struct {
	PartitionID    int32
	LeaderEpoch    int32
	Offset         int64
	LogStartOffset int64
	MaxBytes       int32
}

type FetchPartitionV11List []FetchPartitionV11

func (fetchPartition *FetchPartitionV11) Byte() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, fetchPartition.PartitionID)
	binary.Write(buf, binary.BigEndian, fetchPartition.LeaderEpoch)

	binary.Write(buf, binary.BigEndian, fetchPartition.Offset)
	binary.Write(buf, binary.BigEndian, fetchPartition.LogStartOffset)
	binary.Write(buf, binary.BigEndian, fetchPartition.MaxBytes)
	return buf.Bytes()
}

func (partitions FetchPartitionV11List) Byte() []byte {
	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, uint32(len(partitions)))
	for _, partition := range partitions {
		binary.Write(buf, binary.BigEndian, partition.Byte())
	}
	return buf.Bytes()
}
