// Q: It's called fizz buzz. What you need to do is print the numbers between one and 20.
// However, if the number is divisible by three, say nine, you should print the word fizz instead of the number.
// If the number is divisible by five, say 10, you should print the word buzz instead of the number.
// And if the number is divisible by both three and five, say 15, it should print fizz buzz instead of the number.

package main

import "fmt"

func main() {
	for num := 1; num < 20; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("fizz buzz")
		} else if num%3 == 0 {
			fmt.Println("fizz")
		} else if num%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(num)
		}
	}
}
