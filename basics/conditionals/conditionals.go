package main

import (
	"fmt"
)

func main() {
	// if and else-ifs
	x := -50

	if x > 70 {
		fmt.Println("x is greater")
	} else if x > 40 && x < 60 {
		fmt.Println("x lies in the range 40-60")
	} else if x < 0 {
		fmt.Println("x is negative")
	}

	a, b := 1.4, 2.1
	if fraction := a / b; fraction > 0 {
		fmt.Println("a is more than b")
	}

	// switch
	x = 2

	switch x {
	case 1:
		fmt.Println("hey, it's one")
	case 2:
		fmt.Println("hi there, two")
	default:
		fmt.Println("bye bye")
	}

	// switch without statement
	switch {
	case x == 0:
		fmt.Println("case 0")
	case x > 1:
		fmt.Println("greater than 0")
	case x < 0:
		fmt.Println("less than 0")
	}
}
