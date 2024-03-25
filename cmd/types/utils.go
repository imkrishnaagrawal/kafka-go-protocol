package types

import (
	"encoding/binary"
	"fmt"
	"io"
)

func Uint16bytes(value int) []byte {
	result := make([]byte, 2)
	binary.BigEndian.PutUint16(result, uint16(value))
	return result
}

func Uint32bytes(value int) []byte {
	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, uint32(value))
	return result
}

func Uint32toUint32bytes(value uint32) []byte {
	result := make([]byte, 4)
	binary.BigEndian.PutUint32(result, value)
	return result
}

func Uint64ToUint64bytes(value uint64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, value)
	return result
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
