package main

import "fmt"

func main() {
	var i int = 42
	fmt.Printf("%v,%T", i, i)

	var j float32
	j = float32(i)
	fmt.Println(j)
}
