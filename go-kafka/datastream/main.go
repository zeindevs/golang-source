package main

import (
	"log"

	kafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	topic := "HVSE"
	consumer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "foo",
		"auto.offset.reset": "smallest",
	})
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		ev := consumer.Pool(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("data team reading order:%s\n", string(e.Value))
		case *kafka.Error:
			fmt.Printf("%v\n", e)
		}

	}
}
