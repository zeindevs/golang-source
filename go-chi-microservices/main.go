package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/zeindevs/go-chi-microservices/application"
)

func main() {
	app := application.New(application.LoadConfig())

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
