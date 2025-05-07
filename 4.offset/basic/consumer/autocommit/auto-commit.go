// auto_commit_consumer.go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "offset-demo-topic",
		GroupID: "auto-group", // enables offset commit tracking
		// Auto-commit is enabled by default (every 5s)
	})
	defer reader.Close()

	fmt.Println("Auto Commit Consumer Started")

	for {
		msg, err := reader.ReadMessage(context.Background()) // auto-commit internally
		if err != nil {
			log.Fatalf("Read error: %v", err)
		}

		// Just print — no need to commit manually
		fmt.Printf("AUTO -> Offset: %d, Key: %s, Value: %s\n", msg.Offset, msg.Key, msg.Value)
	}
}
