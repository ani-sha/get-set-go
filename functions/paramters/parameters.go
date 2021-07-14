package main

import "fmt"

func doubleValueAt(values []int, a int) {
	values[a] *= 2
}

func double(n int) {
	n *= 2
}

func doublePtr(n *int) {
	*n *= 2
}

func main() {
	values := []int{2, 5, 7, 8, 9, 10}
	doubleValueAt(values, 4)
	fmt.Println(values)

	n := 10
	// pass by value
	double(n)
	fmt.Println(n)

	// pass by reference
	doublePtr(&n)
	fmt.Println(n)
}
