package types

import (
	"bytes"
	"encoding/binary"
	"net"
)

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
