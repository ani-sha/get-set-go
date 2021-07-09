package main

import (
	"fmt"
)

func main() {
	x, y := 2.0, 10.9
	var mean float64
	mean = (x + y) / 2
	fmt.Println("Mean: ", mean)
}
