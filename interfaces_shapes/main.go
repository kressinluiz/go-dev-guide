package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	t := triangle{height: 3, base: 2}
	sq := square{sideLength: 3}

	printArea(t)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}

func (t triangle) getArea() float64 {
	return t.height * t.base / 2
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
