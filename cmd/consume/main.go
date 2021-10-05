package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"example.com/header_test/internal/types"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	c.SubscribeTopics([]string{"fakeTopic"}, nil)

	for {
		start := time.Now()
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		var t types.Funk
		err = json.Unmarshal(msg.Value, &t)
		if err != nil {
			fmt.Print("failed to unmarshal JSON")
		}
		//if t.Name == "Potato" {
		if msg.Headers[0].Key == "MyService" {
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Process Message: %s\n", string(msg.Value))
			}
			fmt.Printf("%v\n", time.Since(start))
		} else {
			fmt.Println("No messages consumed")
		}
	}

}
