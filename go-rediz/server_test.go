package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/zeindevs/gorediz/client"
)

func TestServerWithMultiClients(t *testing.T) {
	// server := NewServer(Config{})
	// go func() {
	// 	log.Fatal(server.Start())
	// }()
	// time.Sleep(time.Second)

	nClients := 10
	wg := sync.WaitGroup{}
	wg.Add(nClients)
	for i := 0; i < nClients; i++ {
		go func(it int) {
			c, err := client.New("localhost:5001")
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()
			key := fmt.Sprintf("client_%d", it)
			value := fmt.Sprintf("value_%d", it)
			if err := c.Set(context.TODO(), key, value); err != nil {
				log.Fatal(err)
			}
			val, err := c.Get(context.TODO(), key)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("client %d got this val back => %s\n", i, val)
			wg.Done()
		}(i)
	}
	wg.Wait()

	time.Sleep(time.Second)
	if len(server.peers) != 0 {
		t.Fatalf("expected 0 peers but got %d", len(server.peers))
	}
}
