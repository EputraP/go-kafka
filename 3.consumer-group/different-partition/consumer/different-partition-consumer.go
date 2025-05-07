package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// how to run this consumer
//go run consumer.go -name=Consumer1 -group=group-A

func main() {
	var name, group string
	flag.StringVar(&name, "name", "consumer", "consumer instance name")
	flag.StringVar(&group, "group", "group-A", "consumer group ID")
	flag.Parse()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
		GroupID: group,
	})
	defer reader.Close()

	fmt.Printf("%s started (group: %s)\n", name, group)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("read error:", err)
		}
		fmt.Printf("[%s - group %s] partition=%d | key=%s | value=%s\n",
			name, group, m.Partition, string(m.Key), string(m.Value))
	}
}
