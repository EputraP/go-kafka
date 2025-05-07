// manual_commit_consumer.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "offset-demo-topic",
		GroupID:        "manual-group",
		CommitInterval: 0, // disables auto-commit
	})
	defer reader.Close()

	fmt.Println("Manual Commit Consumer Started")

	for {
		msg, err := reader.FetchMessage(context.Background())
		if err != nil {
			log.Fatalf("Fetch error: %v", err)
		}

		// Simulate processing
		fmt.Printf("MANUAL -> Offset: %d, Key: %s, Value: %s\n", msg.Offset, msg.Key, msg.Value)
		time.Sleep(2 * time.Second) // Simulate work

		// Commit only after successful "processing"
		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			log.Printf("Commit failed: %v", err)
		}
	}
}
