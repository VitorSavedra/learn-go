package main

import "fmt"

func printResult(avg float64) {
	if avg >= 7 {
		fmt.Println("Approved:", avg)
	} else {
		fmt.Println("Failed:", avg)
	}
}

func main() {
	printResult(7.5)
	printResult(4.2)
}
