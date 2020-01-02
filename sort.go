package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"d", "r", "o"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
	i := []int{3, 6, 4}
	fmt.Println(i)
	sort.Ints(i)
	fmt.Println(i)

}
