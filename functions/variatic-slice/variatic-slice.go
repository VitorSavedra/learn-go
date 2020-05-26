package main

import "fmt"

func printApproved(approvedList ...string) {
	fmt.Println("Approved list:")
	for i, approved := range approvedList {
		fmt.Printf("%d) %s\n", i+1, approved)
	}
}

func main() {
	approvedList := []string{"Mary", "Pedro", "Raphael", "Mark"}
	printApproved(approvedList...)
}
