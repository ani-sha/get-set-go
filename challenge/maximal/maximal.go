// Q: Collect the maximal value in a slice.

package main

import "fmt"

func main() {
	nums := []int{10, 100, 29, 30, 999, 123}
	max := nums[0]
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Println("Maximal value in the slice is:", max)
}
