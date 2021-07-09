package main

import "fmt"

func main() {
	l := "Hello, Dev Community!"
	fmt.Println("String:" + l)
	fmt.Println(len(l))

	fmt.Printf("l[0] =  %v (type %T)\n", l[0], l[0])

	// slice (start, end)
	fmt.Println(l[4:10])
	// slice (no end)
	fmt.Println(l[4:])
	// slice (no start)
	fmt.Println(l[:4])

}
