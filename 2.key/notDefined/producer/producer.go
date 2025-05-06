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
		Balancer: &kafka.RoundRobin{}, // Default for no key: distributes evenly
	})
	defer writer.Close()

	values := []string{"Message-1", "Message-2", "Message-3"}

	for _, value := range values {
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				// No Key defined
				Value: []byte(value),
			},
		)
		if err != nil {
			log.Printf("could not write message: %v", err)
		} else {
			fmt.Printf("Produced: key=nil, value=%s\n", value)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
