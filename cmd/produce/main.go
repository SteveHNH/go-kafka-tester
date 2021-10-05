package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var s string = `{"name": "Potato", "service": "foo", "subfunk": {"method": "get", "geeze": ["goose","swan","moose"]}}`

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	topic := "fakeTopic"
	for i := 0; i < 1000; i++ {
		p.Produce(&kafka.Message{
			Headers:        []kafka.Header{{Key: "MyService", Value: []byte("foo")}},
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(s),
		}, nil)
	}

	p.Flush(15 * 1000)
}
