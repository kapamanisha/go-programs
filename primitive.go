package main

import (
	"fmt"
)

func main() {
	n := 1 == 1
	m := 1 == 2
	fmt.Println(n)
	fmt.Println(m)
	fmt.Printf("%v,%T\n", n, n)
	var k uint16 = 65
	fmt.Printf("%v,%T\n ", k, k)
	a := 8
	b := 3
	fmt.Println(a | b)
	fmt.Println(a << 4)
	fmt.Println(b >> 3)

}
