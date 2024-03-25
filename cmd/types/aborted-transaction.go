package types

import (
	"bytes"
	"encoding/binary"
)

type AbortedTransaction struct {
	ProducerId  uint64
	FirstOffset uint64
}

func (abortedTransaction *AbortedTransaction) Read(reader *bytes.Reader) {
	binary.Read(reader, binary.BigEndian, &abortedTransaction.ProducerId)
	binary.Read(reader, binary.BigEndian, &abortedTransaction.FirstOffset)
}
