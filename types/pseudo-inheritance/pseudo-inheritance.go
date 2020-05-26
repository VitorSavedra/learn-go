package main

import "fmt"

type car struct {
	name           string
	actualVelocity int
}

type ferrari struct {
	car
	turboEnabled bool
}

func main() {
	f := ferrari{}
	f.name = "F40"
	f.actualVelocity = 0
	f.turboEnabled = true

	fmt.Printf("Does Ferrari %s have the turbo on? %v\n", f.name, f.turboEnabled)
	fmt.Println(f)
}
