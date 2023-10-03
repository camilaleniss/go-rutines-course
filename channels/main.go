package main

import (
	"fmt"
	"time"
)

// What happen here?
func RunProducer() {
	// instantiate the channel
	channel := make(chan string)
	// start the thread
	go runConsumer(channel)
	// first print this
	fmt.Println("Producer Sending Hello")
	// then consumer receive the hello
	channel <- "Hello"
	// Then producer received the bye sent by runConsumer
	fmt.Println("Producer, received", <-channel)
}

func runConsumer(channel chan string) {
	msg := <-channel
	fmt.Println("Consumer, received", msg)
	channel <- "Bye"
}

// What happen here?
func BuffSender() {
	// the buffer channel has capacity for 3 messages
	channel := make(chan string, 3)
	fmt.Println("Sending ONE")
	channel <- "ONE"
	fmt.Println("Sending TWO")
	channel <- "TWO"
	fmt.Println("Sending THREE")
	channel <- "THREE"
	fmt.Println("Done")
	// finish successfully
}

func main() {
	RunProducer()
	time.Sleep(5 * time.Second)
}
