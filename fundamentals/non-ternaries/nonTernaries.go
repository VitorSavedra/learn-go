package main

import "fmt"

// Not exists ternary operators in Go.
func result(avg float64) string {
	if avg >= 6 {
		return "Approved"
	}

	return "Denied"
}

func main() {
	fmt.Println(result(6.2))
}
