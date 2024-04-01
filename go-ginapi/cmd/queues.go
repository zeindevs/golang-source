package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-queue/queue"
	"github.com/golang-queue/queue/core"
)

type jobData struct {
	Name    string
	Message string
}

func (j *jobData) Bytes() []byte {
	fmt.Printf("%s:%s\n", j.Name, j.Message)
	res := sleepSomeTime()
	j = &jobData{Name: "I am awake", Message: res}
	b, _ := json.Marshal(j)
	return b
}

func sleepSomeTime() string {
	seconds := rand.Intn(20)
	sleepTime := time.Duration(seconds) * time.Second
	time.Sleep(sleepTime)
	return fmt.Sprintf("Commander, I slept: %d seconds\n", seconds)
}

func job(i int, rets chan string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		sleepSomeTime()
		rets <- fmt.Sprintf("Hello commander, handle the job: %02d", +i)
		return nil
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	taskN := 100
	rets := make(chan string, taskN)

	q := queue.NewPool(10, queue.WithFn(func(ctx context.Context, m core.QueuedMessage) error {
		v, _ := m.(*jobData)
		json.Unmarshal(m.Bytes(), &v)

		rets <- "Hello, " + v.Name + ", " + v.Message
		return nil
	}))
	defer q.Release()

	for i := 0; i < taskN; i++ {
		go q.QueueTask(job(i, rets))
	}

	for i := 0; i < taskN; i++ {
		fmt.Println("message:", <-rets)
		time.Sleep(20 * time.Millisecond)
	}
}
