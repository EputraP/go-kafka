// group_earliest_consumer.go
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
		GroupID:     "group-earliest",
		StartOffset: kafka.FirstOffset, // equivalent to "auto.offset.reset=earliest"
	})
	defer reader.Close()

	fmt.Println("Consumer group starting from earliest available message")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Read error:", err)
		}
		fmt.Printf("[Earliest] Offset=%d Key=%s Value=%s\n", msg.Offset, msg.Key, msg.Value)
	}
}
