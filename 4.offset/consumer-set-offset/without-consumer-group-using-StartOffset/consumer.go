// start_offset_consumer.go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "offset-demo-topic",
		Partition:   0,
		StartOffset: 5, // Start reading from offset 5
	})
	defer reader.Close()

	fmt.Println("Starting from offset 5 (via StartOffset config)")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Read error:", err)
		}
		fmt.Printf("[StartOffset] Offset=%d Key=%s Value=%s\n", msg.Offset, msg.Key, msg.Value)
	}
}
