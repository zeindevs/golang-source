package main

import (
	"fmt"
	"log"
	"time"

	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type OrderProducer struct {
	producer   *kafka.Producer
	topic      string
	deliverych chan kafka.Event
}

func NewOrderProducer(p *kafka.Producer, topic string) *OrderProducer {
	return &OrderProducer{
		producer:   p,
		topic:      topic,
		deliverych: make(chan kafka.Event, 10000),
	}
}

func (op *OrderProducer) placeOrder(orderType string, size int) error {
	var (
		format  = fmt.Sprintf("%s - %d", orderType, size)
		payload = []byte(format)
	)
	err := op.producer.Produce(&kafka.Message{
		TopicPartiton: kafka.TopicPartiton{
			Topic:     &op.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	}, op.deliverych)
	if err != nil {
		log.Fatal(err)
	}
	<-op.deliverych
	fmt.Printf("placed order on the queue %s\n", format)
	return nil
}

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "foo",
		"acks":              "all",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	op := NewOrderProducer(p, "HVSE")

	for i := 0; i < 1000; i++ {
		if err := op.placeOrder("market", i+1); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 3)
	}
}
