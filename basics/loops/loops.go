package main

import (
	"fmt"
)

func main() {
	// NOTE:
	// For iteration go has only for loops which can be treated as

	// a) Three-component loop (regular for loop)
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// b) While loop
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	// c) For-each range loop
	col := []string{"red", "blue", "black"}
	for i, c := range col {
		fmt.Println(i, c)
	}

	// d) Exit a loop (with/without conditions using continue and break)
	for i := 0; i < 5; i++ {
		if i > 3 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 5; i++ {
		if i < 3 {
			continue
		}
		fmt.Println(i)
	}

	for {
		if a > 3 {
			break
		}
		fmt.Println(a)
		a++
	}

	// e) Infinite loop
	num := 0
	for {
		num++
	}
	fmt.Println(num) // unreachable
}
