package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

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
