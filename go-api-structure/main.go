package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/zeindevs/go-api-structure/api"
	"github.com/zeindevs/go-api-structure/storage"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "The server address")
	flag.Parse()
	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Printf("server up and listening on http://localhost:%s\n", *listenAddr)
	log.Fatal(server.Start())
}
