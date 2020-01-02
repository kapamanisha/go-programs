package main

import "fmt"

func todo() {
	var arr []int
	arr = []int{1, 2, 3, 4}
	arr1 := []string{"manisha"}
	fmt.Println(arr, arr1)
	arr1 = append(arr1, "is", "tall")
	fmt.Println(arr1)
}
func main() {
	todo()

}
