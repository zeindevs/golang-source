package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeindevs/gorediz/client"
)

func TestOfficialRedisClient(t *testing.T) {
	listenAddr := ":5001"
	server := NewServer(Config{
		ListenAddr: listenAddr,
	})
	go func() {
		log.Fatal(server.Start())
	}()
	time.Sleep(time.Millisecond * 400)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", listenAddr),
		Password: "",
		DB:       0,
	})

	testCases := map[string]string{
		"foo":  "bar",
		"a":    "gg",
		"your": "mom",
		"step": "dad",
	}

	for key, val := range testCases {
		if err := rdb.Set(context.Background(), key, val, 0).Err(); err != nil {
			t.Fatal(err)
		}

		newVal, err := rdb.Get(context.Background(), key).Result()
		if err != nil {
			t.Fatal(err)
		}
		if newVal != val {
			t.Fatalf("expected %s but got %s", val, newVal)
		}
	}
}

func TestFooBar(t *testing.T) {
	in := map[string]string{
		"server":  "redis",
		"version": "6.0",
	}
	out := respWriteMap(in)
	fmt.Println(string(out))
}

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
			fmt.Printf("client %d got this val back => %s\n", it, val)
			wg.Done()
		}(i)
	}
	wg.Wait()

	time.Sleep(time.Second)
	if len(server.peers) != 0 {
		t.Fatalf("expected 0 peers but got %d", len(server.peers))
	}
}
