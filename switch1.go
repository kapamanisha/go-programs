package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
		fallthrough
	case 19:
		fmt.Println("Today is 15th. Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
		fallthrough
	default:
		fmt.Println("No information available for that day.")
	}
}
