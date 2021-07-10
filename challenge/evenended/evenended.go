// Q:  An even-ended number is a number whose first and last digits are the same.
// For example, one, 11, 121. How many even-ended numbers are there that are multiplication of four-digit numbers?
// For example, if you multiply 1,001 by 1,011, you'll get a number which is even-ended.
package main

import "fmt"

func main() {
	ctr := 0

	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			product := i * j
			s := fmt.Sprintf("%d", product)
			if s[0] == s[len(s)-1] {
				ctr++
			}
		}
	}
	fmt.Println("No. of even-ended numbers are: ", ctr)
}
