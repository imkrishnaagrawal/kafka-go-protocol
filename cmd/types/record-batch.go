package types

import (
	"bytes"
	"encoding/binary"
)

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
