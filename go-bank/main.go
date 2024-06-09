package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	seed := flag.Bool("seed", false, "seed to db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("seeding the database")
		seedAccounts(store)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
