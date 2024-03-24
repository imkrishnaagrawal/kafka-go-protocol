package main

import (
	"fmt"
)

func main() {

	kafka := NewKafka()
	kafka.Connect("localhost:9092")
	defer kafka.Close()

	fetchResponse, _ := kafka.Fetch("test")
	fmt.Printf("%+v\n", *fetchResponse)
}
