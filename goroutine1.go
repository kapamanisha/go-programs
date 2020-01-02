package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)
	inform := make(chan int)
	go func() {
		messages <- "hiii"
		if messages == 2 {
			inform <- 2
		}
	}()
	select {
	case msg1 := <-messages:
	case inf1 := <-inform:
		fmt.Println("received messages", msg1)
		fmt.Println("received messages", inform)

	}
}
