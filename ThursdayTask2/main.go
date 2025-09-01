package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Trade struct {
	ID     uint `gorm:"primaryKey"`
	Symbol string
	Type   string // BUY or SELL
	Qty    int
	Price  float64
	Time   time.Time
}

var db *gorm.DB

func main() {
	dsn := "host=localhost user=app password=app123 dbname=mutualfund port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	db = conn
	db.AutoMigrate(&Trade{})

	for {
		fmt.Println("\n1. Add Trade\n2. View All Trades\n3. Get Net Position\n4. Exit")
		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			AddTrade()
		case 2:
			ViewAllTrades()
		case 3:
			GetNetPosition()
		case 4:
			fmt.Println("Exited!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func AddTrade() {
	var symbol, tType string
	var qty int
	var price float64

	fmt.Print("Symbol: ")
	fmt.Scan(&symbol)
	fmt.Print("Type (BUY/SELL): ")
	fmt.Scan(&tType)
	fmt.Print("Quantity: ")
	fmt.Scan(&qty)
	fmt.Print("Price: ")
	fmt.Scan(&price)

	symbol = strings.ToUpper(symbol)
	tType = strings.ToUpper(tType)

	if tType != "BUY" && tType != "SELL" {
		fmt.Println("Type must be BUY or SELL")
		return
	}
	if qty <= 0 || price <= 0 {
		fmt.Println("Quantity and Price must be positive")
		return
	}

	trade := Trade{Symbol: symbol, Type: tType, Qty: qty, Price: price, Time: time.Now()}
	if err := db.Create(&trade).Error; err != nil {
		fmt.Println("Error saving trade:", err)
	} else {
		fmt.Println("Trade added with ID:", trade.ID)
	}
}

func ViewAllTrades() {
	var trades []Trade
	db.Find(&trades)

	if len(trades) == 0 {
		fmt.Println("No trades found")
		return
	}

	fmt.Println("\nID\tSymbol\tType\tQty\tPrice\tTime")
	for _, t := range trades {
		fmt.Printf("%d\t%s\t%s\t%d\t%.2f\t%s\n",
			t.ID, t.Symbol, t.Type, t.Qty, t.Price,
			t.Time.Format("2006-01-02 15:04"))
	}
}

func GetNetPosition() {
	type Pos struct {
		Qty, Invest float64
	}
	positions := make(map[string]Pos)

	var trades []Trade
	db.Find(&trades)

	if len(trades) == 0 {
		fmt.Println("No trades to calculate")
		return
	}

	for _, t := range trades {
		p := positions[t.Symbol]
		switch t.Type {
		case "BUY":
			p.Qty += float64(t.Qty)
			p.Invest += float64(t.Qty) * t.Price
		case "SELL":
			p.Qty -= float64(t.Qty)
			p.Invest -= float64(t.Qty) * t.Price
		default:
			fmt.Println("Unknown trade type:", t.Type)
		}
		positions[t.Symbol] = p
	}

	fmt.Println("\nSymbol\tShares\tNet Investment")
	for sym, p := range positions {
		fmt.Printf("%s\t%.0f\t₹%.2f\n", sym, p.Qty, p.Invest)
	}
}
