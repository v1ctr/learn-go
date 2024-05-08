package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	// By default sends and receives block until both the sender and receiver are ready.
	msg := <-messages
	fmt.Println(msg)
}
