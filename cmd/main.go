package main

import (
	"fmt"
	types "kafka-go/cmd/types"
)

func main() {

	kafka := types.NewKafka()
	kafka.Connect("localhost:9092")
	defer kafka.Close()

	fetchResponses, _ := kafka.Fetch("test")
	for _, fetchResponse := range fetchResponses {
		fmt.Printf("%+v\n", fetchResponse)
	}

}
