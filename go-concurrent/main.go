package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) StartAndListen() {
	// you can name your for loop
free:
	for {
		select {
		// block here until someone is sending a message to the channel
		case msg := <-s.msgch:
			fmt.Printf("received message from: %s payload %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Println("the server is doing a gracefull shutdown")
			break free
		default:
		}
	}
}

func sendMessageToServer(msgch chan Message, payload string) {
	fmt.Println("sending message")

	msg := Message{
		From:    "YouBuyOne",
		Payload: payload,
	}

	msgch <- msg
}

func gracefullQuitServer(quitch chan struct{}) {
	close(quitch)
}

func main() {
	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}

	go s.StartAndListen()

	go func() {
		time.Sleep(500 * time.Millisecond)
		sendMessageToServer(s.msgch, "Hello Sailor!")
	}()

	go func() {
		time.Sleep(1000 * time.Millisecond)
		gracefullQuitServer(s.quitch)
	}()

	select {}
}
