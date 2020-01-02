package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan string)

	go func() {
		for {
			select {
			case done <- "manisha":
				fmt.Println(done)

			case s := <-ticker.C:
				fmt.Println("tick at", s)

			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)

}
