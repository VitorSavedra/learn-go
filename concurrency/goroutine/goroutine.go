package main

import (
	"fmt"
	"time"
)

func speak(name, text string, qty int) {
	for i := 0; i < qty; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteration %d)\n", name, text, i+1)
	}
}

func main() {
	// Example 1
	// speak("Vitor", "How are you?", 3)
	// speak("Savedra", "I'm fine!", 1)

	// Example 2
	// go speak("Vitor", "How are you?", 500)
	// go speak("Savedra", "I'm fine!", 500)

	// Example 3
	go speak("Vitor", "How are you?", 10)
	speak("Savedra", "I'm fine!", 5)
}
