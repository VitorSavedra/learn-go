package math

import (
	"fmt"
	"strconv"
)

// Average is the sum of a collection of numbers divided by the
// count of numbers in the collection.
func Average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	avg := total / float64(len(numbers))
	roundedAvg, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", avg), 64)
	return roundedAvg
}
