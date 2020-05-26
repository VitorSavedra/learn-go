package main

import "fmt"

type numberedGrade float64

func (g numberedGrade) between(start, end float64) bool {
	return float64(g) >= start && float64(g) <= end
}

func numberedGradeToLetteredGrade(g numberedGrade) string {
	if g.between(9.0, 10.0) {
		return "A"
	} else if g.between(7.0, 8.99) {
		return "B"
	} else if g.between(5.0, 6.99) {
		return "C"
	} else if g.between(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(numberedGradeToLetteredGrade(9.8))
	fmt.Println(numberedGradeToLetteredGrade(6.9))
	fmt.Println(numberedGradeToLetteredGrade(2.1))
}
