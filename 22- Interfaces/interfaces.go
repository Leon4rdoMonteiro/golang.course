package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type retangle struct {
	height float64
	width  float64
}

func (r retangle) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (math.Pow(c.radius, 2))
}

func showArea(s shape) {
	fmt.Printf("The shape area is %0.2f.\n", s.area())
}

func main() {
	r := retangle{10, 15}

	showArea(r)

	c := circle{20}

	showArea(c)
}
