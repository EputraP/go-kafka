package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "same-partition-topic",
		Balancer: &kafka.Hash{}, // Ensures same key = same partition
	})
	defer writer.Close()

	key := []byte("user-123")
	for i := 1; i <= 10; i++ {
		value := fmt.Sprintf("Message #%d for user-123", i)
		msg := kafka.Message{
			Key:   key,
			Value: []byte(value),
		}

		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			log.Fatal("failed to write:", err)
		}
		fmt.Println("Produced:", value)
		time.Sleep(500 * time.Millisecond)
	}
}
