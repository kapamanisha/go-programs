package main

import "fmt"

var (
	name    string = "manisha"
	number  int    = 46
	section string = "B"
)

func main() {
	fmt.Printf("%v,%T", name, name)
	fmt.Printf("%v,%T", number, number)
	fmt.Printf("%v,%T", section, section)
}
