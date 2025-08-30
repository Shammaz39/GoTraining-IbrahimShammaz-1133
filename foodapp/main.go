package main

import (
	"fmt"
	"foodapp/handlers"
	"foodapp/kafka"
	"foodapp/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello APP..!")

	app := fiber.New()

	dsn := "host=localhost user=fooduser password=foodpass dbname=fooddelivery port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	log.Println("DB connected")

	// Auto-Migrate tables
	err = db.AutoMigrate(&models.Order{}, &models.OrderEvent{})
	if err != nil {
		log.Fatal("Failed to migrate tables!")
	}

	log.Println("Tables migrated successfully")

	// ✅ Init Kafka producer
	kafka.InitProducer()

	app.Get("/", func(c *fiber.Ctx) error {
		log.Println("Fiber is connected ")
		return c.SendString("Hello Fiber!")
	})

	// init handler
	orderHandler := handlers.NewOrderHandler(db)

	// register route directly here
	app.Get("/api/v1/orders/:id", orderHandler.GetOrderByID)
	app.Post("/api/v1/orders", orderHandler.CreateOrder)
	app.Get("/api/v1/orders", orderHandler.GetOrders)

	app.Listen(":3000")
}
