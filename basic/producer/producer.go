package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		// Optional: for debugging
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte("Hello, Kafka3!"),
		},
	)

	if err != nil {
		log.Fatal("failed to write message:", err)
	}

	log.Println("✅ Message sent!")
	time.Sleep(time.Second) // let things flush before exit
}
