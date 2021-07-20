package main

import "fmt"

type Point struct {
	X int
	Y int
}

// Without pointer reciever
func (p Point) Move(x int, y int) {
	p.X += x
	p.Y += y
}

// With pointer reciever
func (p *Point) MoveWithPointer(x int, y int) {
	p.X += x
	p.Y += y
}

func main() {
	p := Point{2, 1}
	p.Move(4, 5)
	fmt.Printf("%v\n", p)

	p.MoveWithPointer(4, 5)
	fmt.Printf("%v\n", p)
}
