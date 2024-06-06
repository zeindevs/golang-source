package main

import (
	"fmt"
	"time"
)

type Server struct {
	users  map[string]string
	userch chan string
	quitch chan struct{}
}

func NewServer() *Server {
	return &Server{
		users:  make(map[string]string),
		userch: make(chan string),
		quitch: make(chan struct{}),
	}
}

func (s *Server) Start() {
	go s.loop()
}

func (s *Server) loop() {
free:
	for {
		select {
		case user := <-s.userch:
			s.addUser(user)
			fmt.Printf("adding new user %s\n", user)
		case <-s.quitch:
			fmt.Printf("server need to quit")
			break free
		default:
		}
	}
}

func (s *Server) addUser(user string) {
	s.users[user] = user
}

func main() {
	userch := make(chan string, 1000)

	for i := 0; i < 10; i++ {
		go func(i int) {
			userch <- fmt.Sprintf("user_%d", i)
		}(i)
	}

	for user := range userch {
		fmt.Println(user)
	}
}

func quitch() {
	server := NewServer()
	server.Start()

	go func() {
		time.Sleep(2 * time.Second)
		close(server.quitch)
	}()

	// this blocks
	select {}
}

func basic() {
	userch := make(chan string, 2)

	userch <- "Bro" // blocking

	userch <- "Zio" // blocking

	userch <- "fooBarBaz" // is waiting for a consumer of the channel (userName)

	user := <-userch

	fmt.Println(user)
}

func sendMessage(msgch <-chan string) {
	msg := <-msgch
	fmt.Println(msg)
}
