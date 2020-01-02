package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("girl:", s)
	s[0] = "apple"
	s[1] = "is"
	s[2] = "fruit"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
}
