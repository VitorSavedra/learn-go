package main

import (
	"fmt"
	"time"
)

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func primes(num int, c chan int) {
	start := 2

	for i := 0; i < num; i++ {
		for prime := start; ; prime++ {
			if isPrime(prime) {
				c <- prime
				start = prime + 1
				time.Sleep(time.Millisecond * 200)
				break
			}
		}
	}

	close(c)
}

func main() {
	c := make(chan int, 30)
	go primes(100, c)

	for primo := range c {
		fmt.Printf("%d ", primo)
	}

	fmt.Println("End")
}
