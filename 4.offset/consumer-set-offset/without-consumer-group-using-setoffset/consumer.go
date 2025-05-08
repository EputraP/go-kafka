// manual_offset_consumer.go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Manual partition mode
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     "offset-demo-topic",
		Partition: 0,
	})
	defer reader.Close()

	// Seek to a specific offset (e.g., offset 2)
	if err := reader.SetOffset(2); err != nil {
		log.Fatal("Failed to set offset:", err)
	}

	fmt.Println("Reading from offset 2 (manual partition)")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Read error:", err)
		}
		fmt.Printf("[Manual] Offset=%d Key=%s Value=%s\n", msg.Offset, msg.Key, msg.Value)
	}
}
