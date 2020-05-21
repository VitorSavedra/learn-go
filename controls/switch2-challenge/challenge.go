package main

import "fmt"

func numberedGradeToLetteredGrade(grade float64) string {
	switch {
	case grade >= 9 && grade <= 10:
		return "A"
	case grade >= 8 && grade < 9:
		return "B"
	case grade >= 5 && grade < 8:
		return "C"
	case grade >= 3 && grade < 5:
		return "D"
	case grade >= 0 && grade < 3:
		return "E"
	default:
		return "Invalid"
	}
}

func main() {
	fmt.Println(numberedGradeToLetteredGrade(10))
	fmt.Println(numberedGradeToLetteredGrade(5))
	fmt.Println(numberedGradeToLetteredGrade(1.4399))
	fmt.Println(numberedGradeToLetteredGrade(-2.8))
	fmt.Println(numberedGradeToLetteredGrade(11))
}
