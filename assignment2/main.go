package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}
type square struct {
	length float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2.0
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sq := square{
		length: 3.14,
	}
	tr := triangle{
		base:   5.1,
		height: 7.3,
	}

	printArea(sq)
	printArea(tr)
}
