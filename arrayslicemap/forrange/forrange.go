package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	for i, number := range numbers {
		fmt.Printf("Index: %d / Element: %d\n", i, number)
	}

	// Ignoring index
	for _, number := range numbers {
		fmt.Println(number)
	}
}
