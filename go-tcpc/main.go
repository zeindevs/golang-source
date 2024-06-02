package main

import (
	"fmt"
	"log"

	"github.com/zeindevs/go-tcpc/tcpc"
)

func main() {
	channel, err := tcpc.New[string](":3000", ":4000")
	if err != nil {
		log.Fatal(err)
	}
	channel.Sendchan <- "GG"

	msg := <-channel.Recvchan

	fmt.Println("received msg from channel (:4000) over TCP: ", msg)
}
