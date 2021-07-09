package main

import (
	"fmt"
)

type Car struct {
	Name  string
	Color string
}

func main() {
	audi := Car{"Audi", "White"}
	fmt.Println(audi)
	fmt.Printf("%+v\n", audi)
}
