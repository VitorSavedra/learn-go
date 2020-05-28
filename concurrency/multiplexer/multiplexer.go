package main

import (
	"fmt"

	"github.com/cod3rcursos/html"
)

func forward(origin <-chan string, target chan string) {
	for {
		target <- <-origin
	}
}

func merge(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go forward(input1, c)
	go forward(input2, c)
	return c
}

func main() {
	c := merge(
		html.Title("https://google.com", "https://youtube.com"),
		html.Title("https://gmail.com", "https://golang.org"),
	)

	fmt.Println(<-c, "|", <-c)
	fmt.Println(<-c, "|", <-c)
}
