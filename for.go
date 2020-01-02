package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Tuesday, time.Wednesday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}
}
