package main

import (
	"fmt"

	"github.com/zeindevs/go-dependency-injection/internal/client"
	"github.com/zeindevs/go-dependency-injection/internal/db"
	"github.com/zeindevs/go-dependency-injection/internal/logger"
)

func main() {
	fmt.Println("Running my cool DI app, let's get ready to play!")

	logger := logger.NewLogger()

	dbService, err := db.NewDbService(logger)
	if err != nil {
		logger.Error(err)
	}

	client, err := client.NewClient(logger, dbService)
	logger.Debug("Creating players...")
	client.AddPlayers()
	logger.Debug("Starting the battle...\n")
	client.PlayRounds(3)
	logger.Debug("Game complete.")
}
