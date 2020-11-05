package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	model "github.com/imhshekhar47/hs-rtta/evt-model"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	log "github.com/sirupsen/logrus"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:29092"})

	if err != nil {
		log.Println("Failed to create producer")
		panic(err)
	}

	defer producer.Close()

	topic := "trades"

	go func() {
		for evt := range producer.Events() {
			switch evtType := evt.(type) {
			case *kafka.Message:
				if evtType.TopicPartition.Error != nil {
					log.Println(fmt.Sprintf("Delivery Failed: %v", evtType.TopicPartition))
				} else {
					log.Println(fmt.Sprintf("Delivery Success: %v", evtType.TopicPartition))
				}
			}
		}
	}()

	//trades := []string{"BUY", "SELL", "SELL", "BUY", "SELL", "BUY", "BUY"}
	trades := []model.TradeCall{
		model.NewBuy("AMZN", rand.Intn(10), 100+rand.Float32()),
	}
	for _, call := range trades {
		json, err := json.Marshal(call)
		if err != nil {
			log.Fatal("Could not serialize to json")
		}
		log.Println(fmt.Sprintf("Sending: %v", json))
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny},
			Value: []byte(json),
		}, nil)
	}

	producer.Flush(15 * 1000)
	log.Println("Closing producer")
}
