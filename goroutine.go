package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}
func main() {
	// fmt.Println("hiiii")
	say("hiii")
	go say("goog moring")
	say("hii")
}
