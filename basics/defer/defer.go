package main

import "fmt"

func cleanup(name string) {
	fmt.Println("Cleaning up ", name)
}

func worker() {
	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("worker")

	defer cleanup("C")
}

func main() {
	worker()
}
