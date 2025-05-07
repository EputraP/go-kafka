// offset_demo_producer.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "offset-demo-topic"
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	for i := 0; i < 5; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("key-%d", i)),
			Value: []byte(fmt.Sprintf("Message #%d", i)),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Fatal("Write failed:", err)
		}
		fmt.Println("Produced:", string(msg.Value))
		time.Sleep(1 * time.Second)
	}
}
