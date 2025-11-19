package main

import (
	"go-asynq/tasks"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	task, err := tasks.NewEmailDeliveryTask(1, "template-id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}

	// Immediately
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// ProcessIn or ProcessAt
	info, err = client.Enqueue(task, asynq.ProcessIn(24*time.Hour))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("scheduled task: id=%s queue=%s", info.ID, info.Queue)

	// MaxRetry, Queue, Timeout, Deadline, Unique
	info, err = client.Enqueue(task, asynq.MaxRetry(3), asynq.Timeout(3*time.Minute))
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
