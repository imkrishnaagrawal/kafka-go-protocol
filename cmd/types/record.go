package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Record struct {
	RecordSize     int64
	Attributes     int8
	TimestampDelta int64
	OffsetDelta    int64
	KeyLength      int64
	Key            string
	ValueLength    int64
	Value          string
	HeaderLength   uint8
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
