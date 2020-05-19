package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// int
	fmt.Println(1, 2, 100)
	fmt.Println("Literal integer is:", reflect.TypeOf(32000))

	// uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("Byte is", reflect.TypeOf(b))

	// int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("The max value of int is", i1)
	fmt.Println("The type of i1 is", reflect.TypeOf(i1))

	// unicode table
	var i2 rune = 'a'
	fmt.Println("Rune is", reflect.TypeOf(i2))
	fmt.Println(i2)

	// float32, float64
	var x float32 = 49.99
	fmt.Println("The type of 'x' is", reflect.TypeOf(x))
	fmt.Println("The type of literal 49.99 is", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("The type of 'bo' is", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Hello, my name is Tux"
	fmt.Println(s1 + "!")
	fmt.Println("The length of 's1' is", len(s1))

	// string with multiple rows
	s2 := `Hello
	my
	name
	is
	Tux`
	fmt.Println("The length of string is", len(s2))

	// char (?)
	char := 'a'
	fmt.Println("The type of char is", reflect.TypeOf(char))
	fmt.Println(char)
}
