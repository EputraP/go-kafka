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
		Balancer: &kafka.Hash{}, // Hash ensures key-based partitioning
	})
	defer writer.Close()

	messages := []kafka.Message{
		{Key: []byte("user-1"), Value: []byte("Created")},
		{Key: []byte("user-2"), Value: []byte("Updated")},
		{Key: []byte("user-3"), Value: []byte("Deleted")},
	}

	for _, msg := range messages {
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Printf("could not write message: %v", err)
		} else {
			fmt.Printf("Produced: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
		}
		time.Sleep(500 * time.Millisecond)
	}
}
