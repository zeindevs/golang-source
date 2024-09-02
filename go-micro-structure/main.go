package main

import (
	"log"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	apiServer := NewApiServer(NewLoggingService(svc))

	log.Println("server up and listening on http://localhost:3000")
	log.Fatal(apiServer.Start(":3000"))
}
