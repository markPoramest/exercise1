package main

import "fmt"

type square struct {
	side float64
}

func (s square) calculates() float64 {
	return s.side * s.side
}

func (s square) info() {
	area := s.calculates()
	fmt.Println("Area of square is: ", area)
}

type circle struct {
	radius float64
}

func (c circle) calculates() float64 {
	return c.radius * c.radius * 3.14
}

func (c circle) info() {
	area := c.calculates()
	fmt.Println("Area of circle is: ", area)
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
	s1.info()
	c1 := circle{
		radius: 5,
	}
	calculateArea(c1)
	c1.info()
}
