package main

import (
	"fmt"
	"strconv" //string conversions
)

func main() {
	var i int = 12
	f := float64(i)
	s := strconv.Itoa(i) //converting int to string
	var big int64 = 100
	var little = int8(big)
	a := 5 / 2

	fmt.Println(i, f, s, little, a)
	fmt.Printf("%q\n", s)
	user := "manu"
	age := 22
	fmt.Println(" the " + user + " age is " + strconv.Itoa(age) + " good")
	// fmt.Sprint used for convert float to string
	b := 2.5
	fmt.Println(" she got " + fmt.Sprint(b) + " grades ")
	s1 := "string"
	a1 := []byte(s1)
	fmt.Println(a1)
}
