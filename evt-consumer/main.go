package main

import (
	"encoding/json"
	"fmt"

	model "github.com/imhshekhar47/hs-rtta/evt-model"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	log "github.com/sirupsen/logrus"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "exchange",
		"auto.offset.reset": "earliest",
	}
	consumer, err := kafka.NewConsumer(config)

	if err != nil {
		panic("Failed to start consumer")
	}

	topics := []string{"trades"}
	consumer.SubscribeTopics(topics, nil)

	for {
		message, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Println(fmt.Sprintf("Error: %v", err))
		} else {
			var call model.TradeCall
			err := json.Unmarshal(message.Value, &call)
			if err != nil {
				log.Fatal("Could not deserialize message")
			}
			log.Println(fmt.Sprintf("Message[%s]: %v", message.TopicPartition, call))
		}
	}
	log.Println("Closing consumer")
	consumer.Close()
}
