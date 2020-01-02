package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	signal := make(chan bool)
	go func() {
		messages <- "hi"
	}()
	select {
	case msg1 := <-messages:
		fmt.Println("received message", msg1)
	default:
		fmt.Println("not received message")
	}

	select {
	case messages <- msg1:
		fmt.Println("sent message", msg1)
	default:
		fmt.Println("not sent message")
	}
	select {
	case msg1 := <-messages:
		fmt.Println("received message", msg1)
	case sig := <-signal:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
