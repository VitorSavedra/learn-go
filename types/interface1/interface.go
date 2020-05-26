package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	firstname string
	lastname  string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.firstname + " " + p.lastname
}

func (p product) toString() string {
	return fmt.Sprintf("%s - $%.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = person{"Mark", "Theodor"}
	fmt.Println(thing.toString())
	print(thing)

	thing = product{"Jeans", 79.79}
	fmt.Println(thing.toString())
	print(thing)

	p2 := product{"Shirts", 29.99}
	print(p2)
}
