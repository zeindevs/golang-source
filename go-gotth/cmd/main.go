package main

import (
	"log"
	"os"

	"github.com/go-gotth/internal/server"
	"github.com/go-gotth/internal/store"
)

var (
	PORT = ":9001"
)

func main() {
	logger := log.New(os.Stdout, "[GOTTH] ", log.LstdFlags)

	logger.Println("creating guests store...")
	guestDb := store.NewGuestStore(logger)
	guestDb.AddGuest(store.Guest{Name: "Guest", Email: "sig@mail.com"})

	srv, err := server.NewServer(logger, PORT, guestDb)
	if err != nil {
		logger.Fatalf("error when creating server: %s", err)
		os.Exit(1)
	}
	if err := srv.Start(); err != nil {
		logger.Fatalf("error when starting server: %s", err)
		os.Exit(1)
	}
}
