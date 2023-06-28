package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {
	tr := triangle{height: 10, base: 5}
	sq := square{sideLength: 4}
	printArea(tr)
	printArea(sq)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return (0.5 * t.height * t.base)
}
