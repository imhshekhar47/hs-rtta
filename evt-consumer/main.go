package main

import (
	"google.golang.org/protobuf/proto"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/imhshekhar47/hs-rtta/common"
	model "github.com/imhshekhar47/hs-rtta/evt-model"

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
		log.Fatalf("Failed to start consumer: %v\n", err)
		panic(err)
	}

	topics := []string{"trades"}
	consumer.SubscribeTopics(topics, nil)

	log.Infoln("Consumer ready")
	for {
		message, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Errorf("Error: %v\n", err)
		} else {
			call := &model.TradeCall{}
			err := proto.Unmarshal([]byte(message.Value), call)
			if err != nil {
				log.Errorf("Error: %v\n", err)
			}
			json, err := common.SerializeToJSON(call)
			if err != nil {
				log.Errorf("Failed to build JSON: %v\n", err)
			}
			log.Infof("Recieved [%s]: %s\n", message.TopicPartition, string(json))

		}
	}
	consumer.Close()
	log.Infoln("Consumer done")
}
