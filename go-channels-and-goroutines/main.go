package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	start := time.Now()
	id := getUserByName("john")
	fmt.Println(id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Message, 2)

	wg.Add(2)

	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}

	log.Println(time.Since(start))
}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	_ = id
	time.Sleep(time.Second * 1)

	ch <- &Message{
		friends: []string{
			"john",
			"jane",
			"joe",
		},
	}
	wg.Done()
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	_ = id
	time.Sleep(time.Second * 2)

	ch <- &Message{
		chats: []string{
			"john",
			"jane",
			"joe",
		},
	}
	wg.Done()
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s-2", name)
}

func leaky() {
	ch := make(chan int)

	go func() {
		msg := <-ch
		log.Println(msg)
	}()
}
