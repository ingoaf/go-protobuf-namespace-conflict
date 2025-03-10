package main

import (
	"fmt"
	"go-protobuf-namespace-conflict/example"
	pbExample "go-protobuf-namespace-conflict/example/proto"
)

func main() {
	msg := &example.SampleMessage{
		Text: "Hello, Proto!",
	}

	msg2 := &pbExample.SampleMessage{
		Text: "Hi2",
	}
	fmt.Println("Generated message:", msg, msg2)
}
