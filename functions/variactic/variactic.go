package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("Average: %.2f", average(1.2, 2.3, 3.4, 4.5))
}
