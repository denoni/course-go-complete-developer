package main

import "log"

type shape interface {
	getArea() float64
}

type triangle struct {
	height, base float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t square) getArea() float64 {
	return t.sideLength * t.sideLength
}

func printArea(s shape) {
	log.Println(s.getArea())
}

func main() {
	t := triangle{base: 10, height: 5}
	s := square{sideLength: 7}

	printArea(t)
	printArea(s)
}
