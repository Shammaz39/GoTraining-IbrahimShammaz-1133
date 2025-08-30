🧩 Assignment 2: Portfolio Tracker Using GORM
Concepts to Use:
gorm ORM for PostgreSQL
Structs and methods
Error handling
Basic CRUD
Problem Statement:
Implement a simple portfolio tracking system that stores buy/sell trades in a PostgreSQL DB using GORM.
Features:
Create a Trade struct
Store trades into DB
Provide a function GetNetPosition() that calculates holdings per symbol
Sample Input:
go
CopyEdit
// Buy 100 INFY at ₹1400
AddTrade("INFY", "BUY", 100, 1400)
// Sell 40 INFY at ₹1410
AddTrade("INFY", "SELL", 40, 1410)
Sample Output of GetNetPosition(): m,m.                                                                                                                                                                          
go
CopyEdit
INFY: 60 shares, Net Investment: ₹84,000
Constraints:
Use GORM models and migration
Should include basic error handling
Store trade timestamp