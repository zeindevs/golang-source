package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/zeindevs/go-htmx-monitor/internal/hardware"
	"nhooyr.io/websocket"
)

type server struct {
	subscriberMessageBuffer int
	mux                     http.ServeMux
	subscribersMu           sync.Mutex
	subscribers             map[*subscriber]struct{}
}

type subscriber struct {
	msgch chan []byte
}

func NewServer() *server {
	s := &server{
		subscriberMessageBuffer: 10,
		subscribers:             make(map[*subscriber]struct{}),
	}
	s.mux.Handle("/", http.FileServer(http.Dir("./htmx")))
	s.mux.HandleFunc("/ws", s.subscribeHandler)
	return s
}

func (s *server) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	err := s.subscribe(r.Context(), w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (s *server) addSubscriber(sub *subscriber) {
	s.subscribersMu.Lock()
	s.subscribers[sub] = struct{}{}
	s.subscribersMu.Unlock()
	fmt.Println("Added subscriber", sub)
}

func (s *server) deleteSubscriber(sub *subscriber) {
	s.subscribersMu.Lock()
	delete(s.subscribers, sub)
	s.subscribersMu.Unlock()
	fmt.Println("Delete subscriber", sub)
}

func (s *server) subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var c *websocket.Conn
	sub := &subscriber{
		msgch: make(chan []byte, s.subscriberMessageBuffer),
	}
	s.addSubscriber(sub)
	defer s.deleteSubscriber(sub)

	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		return err
	}
	defer c.CloseNow()

	ctx = c.CloseRead(ctx)
	for {
		select {
		case msg := <-sub.msgch:
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()

			if err := c.Write(ctx, websocket.MessageText, msg); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (s *server) broadcast(msg []byte) {
	s.subscribersMu.Lock()
	for subscriber := range s.subscribers {
		fmt.Println("Send broadcast to subscriber")
		subscriber.msgch <- msg
	}
	s.subscribersMu.Unlock()
}

func main() {
	log.Println("starting system monitor")
	srv := NewServer()
	go func(s *server) {
		for {
			sysSection, err := hardware.GetSystemSection()
			if err != nil {
				log.Println(err.Error())
			}
			diskSection, err := hardware.GetDiskSection()
			if err != nil {
				log.Println(err.Error())
			}
			cpuSection, err := hardware.GetCpuSection()
			if err != nil {
				log.Println(err.Error())
			}
			timestamp := time.Now().Format("2006-01-2 15:04:05")

			html := []byte(`
        <div hx-swap-oob="innerHTML:#update-timestamp"><span class="h-2 w-2 bg-green-500 rounded-full"></span>` + timestamp + `</div>
        <div hx-swap-oob="innerHTML:#system-data">` + sysSection + `</div>
        <div hx-swap-oob="innerHTML:#cpu-data">` + cpuSection + `</div>
        <div hx-swap-oob="innerHTML:#disk-data">` + diskSection + `</div>
      `)
			s.broadcast(html)
			fmt.Println("get system stats")
			time.Sleep(1 * time.Second)
		}
	}(srv)

	fmt.Println("Server listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", &srv.mux)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
