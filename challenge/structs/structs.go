package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	Center Point
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Length must be > 0")
	}

	s := &Square{Center: Point{x, y}, Length: length}
	return s, nil
}

func (p *Point) Move(x int, y int) {
	p.X += x
	p.Y += y
}

func (s *Square) Move(x int, y int) {
	s.Center.Move(x, y)
}

func (s *Square) Area() int {
	return s.Length * s.Length
}

func main() {
	s, err := NewSquare(1, 3, 10)
	if err != nil {
		log.Fatalf("Can't create square!")
	}
	s.Move(2, 4)
	fmt.Printf("%v\n", s)

	fmt.Println(s.Area())
}
