package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "demo-topic"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topic,
		GroupID:   "demo-group",
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	defer r.Close()

	fmt.Println("Consuming messages...")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("could not read message: %v", err)
		}
		fmt.Printf("Consumed from partition %d: key=%s, value=%s\n", m.Partition, string(m.Key), string(m.Value))
	}
}
