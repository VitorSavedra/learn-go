package main

import "fmt"

func numberedGradeToLetteredGrade(g float64) string {
	var grade = int(g)
	switch grade {
	case 10:
		fallthrough // No break
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 3, 4:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid"
	}
}

func main() {
	fmt.Println(numberedGradeToLetteredGrade(9.8))
	fmt.Println(numberedGradeToLetteredGrade(6.9))
	fmt.Println(numberedGradeToLetteredGrade(11))
	fmt.Println(numberedGradeToLetteredGrade(-1.1))
}
