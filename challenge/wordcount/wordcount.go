// Q: Count how many times each word appears in a text.

package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins 
	Sew me a sail
	To catch me the wind
	`

	words := strings.Fields(text)
	fmt.Println(words)

	counts := map[string]int{} //empty map

	for _, word := range words {
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}
