package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestNewRedisClient(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:5001",
		Password: "",
		DB:       0,
	})
	fmt.Println(rdb)

	if err := rdb.Set(context.Background(), "foo", "bar", 0).Err(); err != nil {
		t.Fatal(err)
	}

	val, err := rdb.Get(context.Background(), "foo").Result()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("go this value", val)
}

func TestNewClient1(t *testing.T) {
	c, err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	if err := c.Set(context.TODO(), "foo", 1); err != nil {
		log.Fatal(err)
	}
	val, err := c.Get(context.TODO(), "foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GET => ", val)
}

func TestNewClient(t *testing.T) {
	c, err := New("localhost:5001")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("SET => ", fmt.Sprintf("foo_%d", i))
		if err := c.Set(context.TODO(), fmt.Sprintf("foo_%d", i),
			fmt.Sprintf("bar_%d", i)); err != nil {
			log.Fatal(err)
		}
		val, err := c.Get(context.TODO(), fmt.Sprintf("foo_%d", i))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("GET => ", val)
	}

}
