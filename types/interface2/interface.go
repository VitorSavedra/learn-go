package main

import "fmt"

type sporty interface {
	enableTurbo()
}

type ferrari struct {
	model          string
	turboEnabled   bool
	actualValocity int
}

func (f *ferrari) enableTurbo() {
	f.turboEnabled = true
}

func main() {
	car1 := ferrari{"F40", false, 0}
	car1.enableTurbo()

	var car2 sporty = &ferrari{"F40", false, 0}
	car2.enableTurbo()

	fmt.Println(car1, car2)
}
