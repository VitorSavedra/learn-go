package main

import "fmt"

func main() {
	approved := make(map[int]string)
	approved[123456789] = "Mary"
	approved[987654321] = "John"
	approved[456456456] = "Ana"
	fmt.Println(approved)

	for document, name := range approved {
		fmt.Printf("Name: %s | Document: %d\n", name, document)
	}

	fmt.Println(approved[123456789])
	delete(approved, 123456789)
	fmt.Println(approved)
}
