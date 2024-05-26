package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zeindevs/microservices-grpc/client"
)

func main() {
	client := client.NewClient("http://localhost:3000")

	price, err := client.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\b", price)
}
