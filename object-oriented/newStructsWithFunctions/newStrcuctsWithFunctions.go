package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

// Create a new trade and validates input
func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol can't be empty")
	}

	if volume <= 0 {
		return nil, fmt.Errorf("Volume must be >= 0 (was %d)", volume)
	}

	if price <= 0.0 {
		return nil, fmt.Errorf("Price must be >= 0 (was %f)", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}

	return trade, nil
}

func main() {
	t1, err := NewTrade("MSW", 20, 99.67, true)
	if err != nil {
		fmt.Println("Error can't create trade -", err)
		os.Exit(1)
	}

	fmt.Println(t1)
}
