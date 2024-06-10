package main

import (
	"flag"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
	"github.com/zeindevs/microservices-hollywood/consumer"
)

func main() {
	accID := flag.Int64("account-id", 0, "id of account")
	flag.Parse()

	e := actor.NewEngineConfig()
	a, err := actor.NewEngine(e)
	if err != nil {
		panic(err)
	}
	r := remote.New("127.0.0.1:40000", remote.Config{})
	e.WithRemote(r)

	// accIDStr := strconv.Itoa(int(*accID))

	a.Spawn(consumer.NewBroker(*accID), "broker", actor.WithInboxSize(10000))

	select {}
}
