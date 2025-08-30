🧩 Assignment 3: Order Executor with Interfaces
Concepts to Use:
Interfaces
Struct embedding
Error handling
Dependency injection
Problem Statement:
Design an order executor system where different order types (MarketOrder, LimitOrder) implement a common interface Order.
Interface:
go
CopyEdit
type Order interface {
   Execute() error
}
Requirements:
Create 2 types: MarketOrder, LimitOrder
Each must have an Execute() method
Implement a ProcessOrder(Order) function to call the interface
Sample Output:
go
CopyEdit
Processing Market Order: Buying 50 INFY at Market Price
Processing Limit Order: Buying 100 AAPL at ₹174.25
Extension:
Add error simulation (e.g., if limit price is below market)
Use fmt.Errorf for rich error reporting
 