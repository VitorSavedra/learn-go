package main

import "fmt"

func main() {
	employeesAndSalaries := map[string]float64{
		"Ana":       11000.90,
		"Mary":      15000.65,
		"Stephanie": 8000.00,
		"Sabrina":   27000.20,
	}

	employeesAndSalaries["Samanta"] = 7500.00

	for name, salary := range employeesAndSalaries {
		fmt.Printf("Name: %s / Salary: %.2f\n", name, salary)
	}
}
