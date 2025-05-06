package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "demo-topic"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.Hash{}, // important: use Hash balancer to honor keys
	})
	defer writer.Close()

	key := "user-123"
	values := []string{"Created", "Updated", "Deleted"}

	for _, value := range values {
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(key),   // Same key
				Value: []byte(value), // Different value
			},
		)
		if err != nil {
			log.Printf("could not write message: %v", err)
		} else {
			fmt.Printf("Produced: key=%s, value=%s\n", key, value)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
