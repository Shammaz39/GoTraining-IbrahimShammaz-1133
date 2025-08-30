package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

// InitProducer initializes the Kafka writer (producer)
func InitProducer() {
	writer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "food-orders",
		Balancer: &kafka.LeastBytes{},
	}
	log.Println("✅ Kafka producer initialized")
}

// Publish sends a message to Kafka
func Publish(key, value string) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
		Time:  time.Now(),
	}

	return writer.WriteMessages(context.Background(), msg)
}
