package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	lines_yesterday := "20"
	lines_today := "50"
	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}
	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday
	fmt.Println(lines_more)
}
