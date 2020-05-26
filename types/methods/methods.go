package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstname string
	lastname  string
}

func (p person) getFullName() string {
	return p.firstname + " " + p.lastname
}

func (p *person) setFullName(fullName string) {
	names := strings.Split(fullName, " ")
	p.firstname = names[0]
	p.lastname = names[1]
}

func main() {
	p1 := person{"Vitor", "Savedra"}
	fmt.Println(p1.getFullName())

	p1.setFullName("Manolo Silva")
	fmt.Println(p1.getFullName())
}
