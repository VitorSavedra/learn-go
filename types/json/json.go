package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {
	// Struct to JSON
	p1 := product{1, "Laptop", 1899.90, []string{"Offer", "Electronic"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// JSON to Struct
	var p2 product
	jsonString := `{"id":2,"name":"Mouse","price":20.9,"tags":["Offer","Eletronic"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[1])
}
