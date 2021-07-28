package main

import (
	"fmt"
	"os"
)

func main() {
	value := []int{2, 4, 7, 8}
	// This will cause a panic
	v := value[10]
	fmt.Println(v)

	file, err := os.Open("no-such-file")
	if err != nil {
		// built-in panic function
		panic(err)
	}
	defer file.Close()
	fmt.Println("File was opened.")
}
