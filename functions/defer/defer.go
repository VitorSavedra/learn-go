package main

import "fmt"

func getApprovedValue(number int) int {
	defer fmt.Println("Done!")
	if number >= 5000 {
		fmt.Println("High value...")
		return 5000
	}

	fmt.Println("Low value...")
	return number
}

func main() {
	fmt.Println(getApprovedValue(6000))
	fmt.Println(getApprovedValue(3000))
}
