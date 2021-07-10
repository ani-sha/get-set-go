package main

import "fmt"

func main() {
	languages := []string{"Go", "Java", "Ruby", "C++"}
	fmt.Printf("languages: %v (type %T)\n", languages, languages)

	fmt.Println(len(languages))
	fmt.Println(languages[1])
	fmt.Println(languages[1:])
	fmt.Println(languages[:2])

	// print using for loop
	for l := 0; l < len(languages); l++ {
		fmt.Println(languages[l])
	}

	//print using single value range
	for l := range languages {
		fmt.Println(languages[l])
	}

	//print using double value range
	for l, name := range languages {
		fmt.Printf("%s at %d position\n", name, l)
	}

	//print using double value range, ignore index
	for _, name := range languages {
		fmt.Println(name)
	}

	// append
	languages = append(languages, "C", "Python")
	fmt.Println(languages)
}
