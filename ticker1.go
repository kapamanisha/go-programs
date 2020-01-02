package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	message := make(chan int)
	go func() {
		message <- "4"

		for {
			select {
			case <-message:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()
	time.Sleep(2000 * time.Millisecond)
	ticker.Stop()
	message <- "manu"
	fmt.Println("ticker stopped")
}
