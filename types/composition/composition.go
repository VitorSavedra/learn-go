package main

import "fmt"

type sporty interface {
	enableTurbo()
}

type luxurious interface {
	autoParking()
}

type sportyLuxurious interface {
	sporty
	luxurious
}

type bmw7 struct{}

func (b bmw7) enableTurbo() {
	fmt.Println("Turbo ON!")
}

func (b bmw7) autoParking() {
	fmt.Println("Parking...")
}

func main() {
	var b sportyLuxurious = bmw7{}
	b.enableTurbo()
	b.autoParking()
}
