//Interfaces are collection of methods.
// Each type implements all the methods in an interface.

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Square struct {
	Length float64
}

type Circle struct {
	Radius float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func main() {
	s := &Square{20}
	fmt.Println(s.Area())

	c := &Circle{14}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)
}
