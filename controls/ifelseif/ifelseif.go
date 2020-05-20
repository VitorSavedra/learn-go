package main

import "fmt"

func numberedGradeToLetteredGrade(grade float64) string {
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 8 && grade < 9 {
		return "B"
	} else if grade >= 5 && grade < 8 {
		return "C"
	} else if grade >= 3 && grade < 5 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(numberedGradeToLetteredGrade(8))
	fmt.Println(numberedGradeToLetteredGrade(6.9))
	fmt.Println(numberedGradeToLetteredGrade(2.1))
}
