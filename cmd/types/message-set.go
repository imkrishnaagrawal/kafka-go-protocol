package types

import (
	"bytes"
	"encoding/binary"
)

type MessageSet struct {
	RecordBatch RecordBatch
}

func (messageSet *MessageSet) Read(messageReader *bytes.Reader) {
	messageSet.RecordBatch = RecordBatch{}
	for {
		len := messageReader.Len()
		if len <= 0 {
			break
		}
		recordBatch := RecordBatch{}
		recordBatch.Read(messageReader)
		binary.Read(messageReader, binary.BigEndian, &recordBatch.RecordLength)

		record := Record{}
		record.Read(messageReader)
		binary.Read(messageReader, binary.BigEndian, record.HeaderLength)
		for h := 0; h < int(record.HeaderLength); h++ {
			header := Header{}
			header.Read(messageReader)
			record.Headers = append(record.Headers, header)
		}
		recordBatch.Records = append(recordBatch.Records, record)

		messageSet.RecordBatch = recordBatch
	}
}
