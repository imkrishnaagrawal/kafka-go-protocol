package main

import (
	"fmt"
)

func main() {

	kafka := NewKafka()
	kafka.Connect("localhost:9092")
	defer kafka.Close()

	fetchResponses, _ := kafka.Fetch("test")
	for _, fetchResponse := range fetchResponses {
		fmt.Printf("%+v\n", fetchResponse)
	}

}
