package main

import (
	"fmt"
	"time"
)

func checkType(i interface{}) string {
	switch i.(type) {
	case int:
		return "Integer"
	case float32, float64:
		return "Float"
	case string:
		return "String"
	case func():
		return "Function"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println(checkType(2.3))
	fmt.Println(checkType(1))
	fmt.Println(checkType("Lorem"))
	fmt.Println(checkType(func() {}))
	fmt.Println(checkType(time.Now()))
}
