package main

import (
	"math/rand"
	"time"

	model "github.com/imhshekhar47/hs-rtta/evt-model"
	"google.golang.org/protobuf/proto"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	log "github.com/sirupsen/logrus"
)

var (
	r      = rand.New(rand.NewSource(time.Now().UnixNano()))
	stocks = []string{"AMZN", "MSFT", "GOGL", "APPL", "TSLA", "LCMR", "WMRT"}
)

func rInt(min int, max int) int32 {
	return int32(min + r.Intn(max-min))
}

func rFloat(min float32, max float32) float32 {
	return min + float32(int(max-min)) + r.Float32()
}

func rStock() string {
	return stocks[rInt(0, len(stocks))]
}

func rTradeCall() *model.TradeCall {
	stock := rStock()
	action := model.CallType_BUY
	if rInt(0, 1) == 0 {
		action = model.CallType_SELL
	}
	return &model.TradeCall{
		Action: action,
		Stock:  stock,
		Units:  rInt(1, 20),
		Price:  rFloat(10, 20),
	}
}

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	}
	producer, err := kafka.NewProducer(config)

	if err != nil {
		log.Fatalf("Failed to create producer: %v\n", err)
		panic(err)
	}

	defer producer.Close()

	topic := "trades"

	go func() {
		for evt := range producer.Events() {
			switch evtType := evt.(type) {
			case *kafka.Message:
				if evtType.TopicPartition.Error != nil {
					log.Infof("Failed: %v\n", evtType.TopicPartition)
				} else {
					log.Infof("Delivered: %v\n", evtType.TopicPartition)
				}
			}
		}
	}()

	//trades := []string{"BUY", "SELL", "SELL", "BUY", "SELL", "BUY", "BUY"}

	log.Infoln("Producer ready")
	for i := 0; i < 100; i++ {
		call := rTradeCall()
		bytes, err := proto.Marshal(call)
		if err != nil {
			log.Fatalf("Error: %v\n", err)
		}
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny},
			Value: bytes,
		}, nil)
	}

	producer.Flush(15 * 1000)
	log.Infoln("Producer done")
}
