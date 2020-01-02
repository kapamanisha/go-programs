package main

import (
	"fmt"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	breadth, length float64
}

func (r rect) area() float64 {
	return  Math.r.breadth * r.length
}
func (r rect) perim() float64 {
	return (2*r.breadth +  2*r.length
}
func measure(g geomertry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
func main() {
	r := rect{breadth: 3, length: 5}
	fmt.Println("r.area")
	measure(r)
}
