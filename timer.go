package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println(timer1)
	go func() {
		<-timer1.C
		fmt.Println("timer1 experied")
	}()
	stop1 := timer1.Stop()
	if stop1 {
		fmt.Println("timer1 stopped")
	}
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 experied")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}
}
