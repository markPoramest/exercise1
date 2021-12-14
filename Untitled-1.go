package main

import "fmt"

type square struct {
	side float64
}

func (s square) calculates() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) calculates() float64 {
	return c.radius * c.radius * 3.14
}

type shape interface {
	calculates() float64
}

func calculateArea(s shape) {
	fmt.Println(s.calculates())
}
func main() {
	s1 := square{
		side: 5,
	}
	calculateArea(s1)
	c1 := circle{
		radius: 5,
	}
	calculateArea(c1)
}
