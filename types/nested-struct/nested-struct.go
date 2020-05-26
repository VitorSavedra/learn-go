package main

import "fmt"

type item struct {
	productID int
	quantity  int
	price     float64
}

type order struct {
	userID int
	itens  []item
}

func (o order) totalValue() float64 {
	total := 0.0
	for _, item := range o.itens {
		total += item.price * float64(item.quantity)
	}

	return total
}

func main() {
	order := order{
		userID: 1,
		itens: []item{
			item{productID: 1, quantity: 2, price: 12.10},
			item{productID: 2, quantity: 1, price: 23.49},
			item{productID: 3, quantity: 10, price: 3.00},
		},
	}

	fmt.Printf("Total value: $%.2f", order.totalValue())
}
