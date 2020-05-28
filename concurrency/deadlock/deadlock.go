package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // Blocking operation
	fmt.Println("After...")
}

func main() {
	c := make(chan int) // Without buffer
	go routine(c)

	fmt.Println(<-c) // Blocking operation
	fmt.Println("Read")
	fmt.Println(<-c) // Deadlock
	fmt.Println("End")

	// Deadlock return:
	// fatal error: all goroutines are asleep - deadlock!
}
