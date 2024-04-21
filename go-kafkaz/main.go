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

const (
	KeyAPIVersions = 18
)

type ApiVersion struct {
	APIKey     int16
	MinVersion int16
	MaxVersion int16
}

type ApiVersionResponse struct {
	ErrorCode      int16
	ApiVersions    []ApiVersion
	ThrottleTimeMs int32
}

func (avr *ApiVersionResponse) Encode(w io.Writer) error {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, int32(0))  // placeholder
	binary.Write(buf, binary.BigEndian, int32(18)) // apiKEY version...
	binary.Write(buf, binary.BigEndian, int32(3))  // apiVersion
	binary.Write(buf, binary.BigEndian, avr.ErrorCode)
	binary.Write(buf, binary.BigEndian, len(avr.ApiVersions))

	for _, version := range avr.ApiVersions {
		binary.Write(buf, binary.BigEndian, version.APIKey)
		binary.Write(buf, binary.BigEndian, version.MinVersion)
		binary.Write(buf, binary.BigEndian, version.MaxVersion)
	}

	binary.Write(buf, binary.BigEndian, avr.ThrottleTimeMs)
	binary.BigEndian.PutUint32(buf.Bytes(), uint32(buf.Len()-4))
	return nil
}

// HeaderHeader
type Header struct {
	Size int32

	APIKey        int16
	APIVersion    int16
	CorrelationID int32
	ClientID      string
}

func (h Header) Encode(w io.Writer) error {
	binary.Write(w, binary.BigEndian, h.Size)
	return nil
}

type APIVersion struct {
	CorrelationID         int32
	ClientID              string
	ClientSoftwareName    string
	ClientSoftwareVersion string
}

func readAPIVersion(r io.Reader) APIVersion {
	var version APIVersion
	binary.Read(r, binary.BigEndian, &version.CorrelationID)

	var size int16
	binary.Read(r, binary.BigEndian, &size)

	clientID := make([]byte, size)
	binary.Read(r, binary.BigEndian, &clientID)

	binary.Read(r, binary.BigEndian, &size)
	clientSoftwareName := make([]byte, size)
	binary.Read(r, binary.BigEndian, &clientSoftwareName)

	clientSoftwareVersion, _ := io.ReadAll(r)

	return APIVersion{
		ClientID:              string(clientID),
		ClientSoftwareName:    string(clientSoftwareName),
		ClientSoftwareVersion: string(clientSoftwareVersion),
	}
}

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
		rawMsg := buf[:n]
		r := bytes.NewReader(rawMsg)
		var header Header
		binary.Read(r, binary.BigEndian, &header)

		switch header.APIKey {
		case KeyAPIVersions:
			version := readAPIVersion(r)
			slog.Info("server received message from client", "message", version)
			resp := ApiVersionResponse{
				ErrorCode: 0,
				ApiVersions: []ApiVersion{
					{
						APIKey:     0,
						MinVersion: 0,
						MaxVersion: 10,
					},
					{
						APIKey:     KeyAPIVersions,
						MinVersion: 0,
						MaxVersion: 3,
					},
					{
						APIKey:     19,
						MinVersion: 0,
						MaxVersion: 7,
					},
				},
				ThrottleTimeMs: 100,
			}
			fmt.Println("before encoding version")
			resp.Encode(conn)
			fmt.Println("after encoding version")
		default:
			fmt.Println("unhandled message from the client => ", header.APIKey)
		}
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

func readVarint(r io.ByteReader) (int, error) {
	var result int
	var shift uint
	for {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		result |= int(b&0x7f) << shift
		if (b & 0x80) == 0 {
			break
		}
		shift += 7
	}
	return result, nil
}
