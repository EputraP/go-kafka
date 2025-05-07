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
		Topic:    "my-topic",
		Balancer: &kafka.Hash{}, // key-based partitioning
	})
	defer writer.Close()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%d", i%5) // 5 different keys
		value := fmt.Sprintf("message-%d", i)

		err := writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(key),
			Value: []byte(value),
		})
		if err != nil {
			log.Println("failed to write:", err)
		} else {
			fmt.Printf("Produced: key=%s, value=%s\n", key, value)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
