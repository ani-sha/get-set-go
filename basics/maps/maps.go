package main

import "fmt"

func main() {
	stocks := map[string]float64{
		"Amazon": 1209.09,
		"Google": 1699.3,
		"Zomato": 826.00,
	}

	fmt.Println(len(stocks))

	fmt.Println(stocks["Zomato"])

	fmt.Println(stocks["Apple"]) // 0 if key not found

	value, ok := stocks["Apple"]
	if !ok {
		fmt.Println("Apple stocks are not listed.")
	} else {
		fmt.Println(value)
	}

	// Add a new value
	stocks["Apple"] = 1561.72
	fmt.Println(stocks)

	// Delete using a key
	delete(stocks, "Amazon")
	fmt.Println(stocks)

	for key := range stocks {
		fmt.Println(key)
	}

	for key, value := range stocks {
		fmt.Printf("%s ---> %.2f\n", key, value)
	}

}
