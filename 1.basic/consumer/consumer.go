package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
	})

	defer reader.Close()

	log.Println("🔁 Listening for messages...")

	ctx := context.Background()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal("error reading message:", err)
		}
		log.Printf("📩 Message: key=%s value=%s\n", string(msg.Key), string(msg.Value))
	}
}
