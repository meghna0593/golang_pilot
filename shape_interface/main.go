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

// Program that creates two custom struct types 'Triangle' and 'Square' with a 'Shape' interface
func main() {
	ts := triangle{height: 5.0, base: 6.0}
	ss := square{sideLength: 4.0}

	printArea(ts)
	printArea(ss)

}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
