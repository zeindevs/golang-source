package main

import (
	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
	"github.com/zeindevs/microservices-hollywood/consumer"
)

func main() {
	e := actor.NewEngineConfig()
	a, err := actor.NewEngine(e)
	if err != nil {
		panic(err)
	}
	r := remote.New("127.0.0.1:40000", remote.Config{})
	e.WithRemote(r)

	a.Spawn(consumer.NewProcessor(), "broker", actor.WithInboxSize(10000))

	select {}
}
