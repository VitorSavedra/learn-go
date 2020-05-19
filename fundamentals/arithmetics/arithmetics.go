package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum:", a+b)
	fmt.Println("Sub:", a-b)
	fmt.Println("Div:", a/b)
	fmt.Println("Mul:", a*b)
	fmt.Println("Mod:", a%b)
	fmt.Println("And:", a&b) // 11 & 10 = 10
	fmt.Println("Or:", a|b)  // 11 | 10 = 11
	fmt.Println("Xor:", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	fmt.Println("Bigger:", math.Max(c, d))
	fmt.Println("Smaller:", math.Min(c, d))
	fmt.Println("Expo:", math.Pow(c, d))
}
