package main

import (
	"fmt"
	"log"

	"github.com/zeindevs/go-tcpc/tcpc"
)

func main() {
	channel, err := tcpc.New[string](":4000", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	msg := <-channel.Recvchan

	fmt.Printf("received msg from channel (%s) over TCP: (%s)\n", ":3000", msg)

	channel.Sendchan <- msg
}
