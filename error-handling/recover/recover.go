package main

import "fmt"

func safeValue(vals []int, index int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %s\n", err)
		}
	}()
	return vals[index]
}

func main() {
	s := safeValue([]int{2, 3, 4, 8}, 10)
	fmt.Println(s)
}
