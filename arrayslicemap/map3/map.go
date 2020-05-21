package main

import "fmt"

func main() {
	employeesByLetter := map[string]map[string]float64{
		"G": {
			"Gabrielle": 15444.00,
			"Giovanna":  8488.20,
			"Glau":      2700.99,
			"Gizele":    8000.00,
		},
		"M": {
			"Mary":   20000.00,
			"Marina": 11100.20,
		},
		"A": {
			"Anabelle": 5760.00,
			"Ana":      8000.00,
		},
	}

	delete(employeesByLetter, "M")

	for letter, employee := range employeesByLetter {
		fmt.Println(letter, employee)
	}
}
