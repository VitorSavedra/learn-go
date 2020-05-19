package main

import "fmt"

func main() {
	fmt.Print("Same")
	fmt.Print(" row.\n")

	fmt.Println("New")
	fmt.Println("Row")

	x := 3.141516

	xs := fmt.Sprint(x)
	fmt.Println("The value of 'x' is: " + xs)
	fmt.Println("The value of 'x' is:", x)

	fmt.Printf("The value of 'x' is: %.2f.", x)

	a := 1
	b := 1.9999
	c := false
	d := "Ops"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
