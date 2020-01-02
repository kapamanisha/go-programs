package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{"manisha", 22}
	fmt.Println(p)
}
