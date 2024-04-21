package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Messsage struct {
	data []byte
	// ...
}

type Config struct {
	ListenAddr string
}

type Server struct {
	listenAddr string
	ln         net.Listener

	coffsets map[string]int
	buffer   []Messsage
}

func NewServer(cfg Config) *Server {
	return &Server{
		listenAddr: cfg.ListenAddr,
		coffsets:   make(map[string]int),
		buffer:     []Messsage{},
	}
}

func (s *Server) Listen() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}

	slog.Info("server listening on ", "port", s.listenAddr)

	s.ln = ln

	return s.loop()
}

func (s *Server) loop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			if err == io.EOF {
				return err
			}
			slog.Error("server accept error", "err", err)
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	fmt.Println("new connection", conn.RemoteAddr())
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			slog.Error("connection read error", "err", err)
			return
		}
		msg := buf[:n]
		fmt.Println(string(msg))
		r := bytes.NewReader(msg)

		acks := make([]byte, 2)
		binary.Read(r, binary.BigEndian, acks)

		timeoutms := binary.BigEndian.Uint32(msg[2:6])

		fmt.Println("acks: ", acks)
		fmt.Println("timeoutms:", timeoutms)
	}
}

func main() {
	server := NewServer(Config{
		ListenAddr: ":9092",
	})
	go func() {
		log.Fatal(server.Listen())
	}()
	time.Sleep(time.Second * 2)

	fmt.Println("producing...")

	produce()

	time.Sleep(time.Second * 2)
}

func produce() error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	})
	if err != nil {
		return err
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

	topic := "fooTopic"
	for _, word := range []string{"Welcome", "to", "Jungle"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	p.Flush(15 * 1000)

	return nil
}

func consume() error {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return err
	}

	c.SubscribeTopics([]string{"fooTopic"}, nil)

	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if err.(kafka.Error).IsFatal() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	return c.Close()
}
