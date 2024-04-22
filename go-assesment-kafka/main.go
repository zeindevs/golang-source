package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Storer interface {
	Put(MessageState, []byte) error
	Get(MessageState) ([][]byte, error)
}

type Storage struct {
	mu   sync.RWMutex
	data map[MessageState][][]byte
}

func NewStorage() *Storage {
	data := map[MessageState][][]byte{}
	data[MessageStateCompleted] = [][]byte{}
	data[MessageStateFailed] = [][]byte{}
	data[MessageStateProgress] = [][]byte{}
	return &Storage{
		data: data,
	}
}

func (s *Storage) Put(state MessageState, val []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[state] = append(s.data[state], val)
	return nil
}

func (s *Storage) Get(state MessageState) ([][]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[state]
	if !ok {
		return nil, fmt.Errorf("value not found man")
	}
	return val, nil
}

const (
	lenMessage = 1000
)

var (
	topic = "foobartopic"
)

type MessageState int

const (
	MessageStateCompleted MessageState = iota
	MessageStateProgress
	MessageStateFailed
)

type Message struct {
	State MessageState
}

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	produce(cancel)
	c, err := NewConsumer(NewStorage())
	if err != nil {
		log.Fatal(err)
	}
	c.consumeLoop(ctx)
	fmt.Printf("%+v\n", c.Storage.(*Storage).data)
}

type Consumer struct {
	consumer *kafka.Consumer
	Storage  Storer
	quitch   chan struct{}
}

func NewConsumer(storage Storer) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        "localhost:9093",
		"broker.address.family":    "v4",
		"group.id":                 "group1",
		"session.timeout.ms":       6000,
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": false,
	})
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		Storage:  NewStorage(),
		consumer: c,
		quitch:   make(chan struct{}),
	}, nil
}

func (c *Consumer) consumeLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ev := c.consumer.Poll(100)
			if ev != nil {
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				_, err := c.consumer.StoreMessage(e)
				if err != nil {
					fmt.Println("store msg error:", err)
				}
				var msg Message
				if err := json.Unmarshal(e.Value, &msg); err != nil {
					log.Fatal(err)
				}
				if err := c.Storage.Put(msg.State, e.Value); err != nil {
					log.Fatal(err)
				}
			case kafka.Error:
				if e.Code() == kafka.ErrAllBrokersDown {
					break
				}
			}
		}
	}
}

func produce(cancel context.CancelFunc) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9093",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	slog.Info("start producing", "topic", topic, "message", lenMessage)
	for i := 0; i < lenMessage; i++ {
		msg := Message{
			State: MessageState(rand.Intn(3)),
		}

		b, err := json.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: b,
		}, nil)

		if err != nil {
			log.Fatal(err)
		}
	}
	cancel()
	slog.Info("done producing", "topic", topic, "message", lenMessage)
}
