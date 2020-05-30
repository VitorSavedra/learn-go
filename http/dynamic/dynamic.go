package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func rightTime(w http.ResponseWriter, r *http.Request) {
	hour := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Right Hour: %s</h1>", hour)
}

func main() {
	http.HandleFunc("/hour", rightTime)
	log.Println("Running...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
