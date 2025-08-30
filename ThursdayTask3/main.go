package main

import (
	"fmt"
)

type Order interface {
	Execute() error
}

type MarketOrder struct {
	Quantity int
	Symbol   string
}

func (m MarketOrder) Execute() error {
	fmt.Printf("Processing Market Order: %d %s at Market Price\n", m.Quantity, m.Symbol)
	return nil
}

type LimitOrder struct {
	Quantity    int
	Symbol      string
	LimitPrice  float64
	MarketPrice float64
}

func (l LimitOrder) Execute() error {

	if l.LimitPrice < l.MarketPrice {
		return fmt.Errorf("limit price %.2f too low for BUY (market is %.2f)", l.LimitPrice, l.MarketPrice)
	} else {

		fmt.Printf("Processing Limit Order:  %d %s at ₹%.2f\n", l.Quantity, l.Symbol, l.LimitPrice)
	}
	return nil
}

func ProcessOrder(o Order) {
	if err := o.Execute(); err != nil {
		fmt.Println(" Order Failed:", err)
	}
}

func main() {

	var orderType string
	fmt.Print("Enter order type (M or L): ")
	fmt.Scanln(&orderType)

	switch orderType {
	case "M":
		var symbol string
		var qty int

		fmt.Print("Enter stock symbol: ")
		fmt.Scanln(&symbol)

		fmt.Print("Enter quantity: ")
		fmt.Scanln(&qty)

		order := MarketOrder{Symbol: symbol, Quantity: qty}
		ProcessOrder(order)

	case "L":
		var symbol string
		var qty int
		var limitPrice, marketPrice float64

		fmt.Print("Enter stock symbol: ")
		fmt.Scanln(&symbol)

		fmt.Print("Enter quantity: ")
		fmt.Scanln(&qty)

		fmt.Print("Enter limit price: ")
		fmt.Scanln(&limitPrice)

		fmt.Print("Enter market price: ")
		fmt.Scanln(&marketPrice)

		order := LimitOrder{Symbol: symbol, Quantity: qty, LimitPrice: limitPrice, MarketPrice: marketPrice}
		ProcessOrder(order)

	default:
		fmt.Println("Invalid order type! Please type 'M' or 'L'.")
	}
}
