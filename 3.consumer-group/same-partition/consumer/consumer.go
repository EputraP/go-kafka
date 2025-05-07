package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

//how to run this consumer
// go run same_partition_consumer.go -name=Consumer1
// go run same_partition_consumer.go -name=Consumer2

func main() {
	var name string
	flag.StringVar(&name, "name", "consumer", "consumer name")
	flag.Parse()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "same-partition-topic",
		GroupID: "same-key-group",
	})
	defer reader.Close()

	fmt.Printf("%s started\n", name)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("read error:", err)
		}
		fmt.Printf("[%s] Partition=%d | Offset=%d | Key=%s | Value=%s\n",
			name, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
