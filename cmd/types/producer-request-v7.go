package types

type ProduceRequestPartitionV7List []ProduceRequestPartitionV7

type ProduceRequestTopicV7List []ProduceRequestTopicV7

type ProduceRequestV7 struct {
	RequestHeader         RequestHeader
	TransactionalIdLength int16
	TransactionalID       *string
	RequiredAcks          int16
	Timeout               int32
	Topics                ProduceRequestTopicV7List
}

type ProduceRequestTopicV7 struct {
	TopicName  string
	Partitions ProduceRequestPartitionV7List
}

type ProduceRequestPartitionV7 struct {
	PartitionID int32
	MessageSet  MessageSet
}

func (kafka *Kafka) NewProduceRequestV7(topic string, offset int64, correlationID int32) *ProduceRequestV7 {
	return &ProduceRequestV7{
		RequestHeader: RequestHeader{
			APIKey:        API_KEY_PRODUCE,
			APIVersion:    7,
			CorrelationID: correlationID,
			ClientId:      kafka.ClientId,
		},
		TransactionalIdLength: -1,
		TransactionalID:       nil,
		RequiredAcks:          -1,
		Timeout:               30000,
		Topics: ProduceRequestTopicV7List{
			ProduceRequestTopicV7{
				TopicName: topic,
				Partitions: ProduceRequestPartitionV7List{
					{
						PartitionID: 0,
						MessageSet: MessageSet{
							RecordBatch: RecordBatch{
								BaseOffset:           0,
								BatchLength:          64,
								PartitionLeaderEpoch: 0,
								MagicByte:            2,
								CRC:                  4002808660,
								LastOffsetDelta:      0,
								BaseTimestamp:        1711372375519,
								MaxTimestamp:         1711372375519,
								ProducerID:           -1,
								ProducerEpoch:        -1,
								BaseSequence:         -1,
								RecordLength:         1,
								Records: []Record{
									{
										RecordSize:     14,
										Attributes:     0,
										TimestampDelta: 0,
										OffsetDelta:    0,
										KeyLength:      3,
										Key:            "Key",
										ValueLength:    5,
										Value:          "Value",
										HeaderLength:   0,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
