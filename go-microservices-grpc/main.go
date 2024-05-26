package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/zeindevs/microservices-grpc/client"
	"github.com/zeindevs/microservices-grpc/proto"
)

func main() {
	var (
		jsonAddr = flag.String("jsonAddr", ":3000", "http listen address the service is running")
		grpcAddr = flag.String("grpcAddr", ":4000", "grpc listen address the service is running")
		svc      = NewLoggingService(NewMetricsService(&priceService{}))
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			time.Sleep(3 * time.Second)
			resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%+v\n", resp)
		}
	}()

	go makeGRPCServerAndRun(*grpcAddr, svc)
	log.Println("GRPC Server up and listening on port :4000")

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	log.Println("HTTP Server up and listening on port :3000")
	jsonServer.Run()
}
