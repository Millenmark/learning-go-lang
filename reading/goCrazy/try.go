package main

import (
	"fmt"
)

type Message struct {
	Hello string
}

func main() {
	// fmt.Printf("Go version: %s\n", runtime.Version())
	h := Message{Hello: "WORLD"}
	fmt.Printf("%s\n", h)
	fmt.Printf("%+v\n", h)
}