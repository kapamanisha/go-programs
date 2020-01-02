package main

import (
	"fmt"
)

func main() {
	twoD := make([][]int, 5)
	for i := 0; i < 5; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
