package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // Send data to channel (write)
	<-ch    // Get data channel (read)

	ch <- 2
	fmt.Println(<-ch)
}
