package main

import (
	"errors"
	"fmt"
)

func boom() error {
	return errors.New("bangles")
}

func main() {
	err := boom()
	if err != nil {
		fmt.Println("an error occured:", err)
		return
	}
	fmt.Println("anchors awaay !")
	
}
