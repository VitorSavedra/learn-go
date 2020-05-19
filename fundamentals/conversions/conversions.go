package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	// int to string
	grade := 6.9
	finalGrade := int(grade)
	fmt.Println(finalGrade)
	fmt.Println("Final grade:", string(finalGrade))       // Wrong
	fmt.Println("Final grade:", strconv.Itoa(finalGrade)) // Correct

	// string to int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// bool to string
	b, _ := strconv.ParseBool("True")
	if b {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
