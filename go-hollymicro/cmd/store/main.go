package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"reflect"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
	"github.com/zeindevs/go-hollymicro/types"
)

type store struct{}

func newStore() actor.Receiver {
	return &store{}
}

func (s *store) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case *types.CatFact:
		slog.Debug("stored fact into the db", "data", msg.Fact)
	case *actor.PID:
		slog.Info("store got pid", "pid", msg)
	case actor.Started:
		slog.Debug("store is started")
	case actor.Stopped:
		slog.Debug("store is stopped")
	default:
		slog.Warn("store got unknown message", "data", msg, "type", reflect.TypeOf(msg).String())
	}
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

	listenAddr := flag.String("listenAddr", "127.0.0.1:4000", "todo")
	flag.Parse()

	r := remote.New(*listenAddr, remote.NewConfig())
	e, err := actor.NewEngine(actor.NewEngineConfig().WithRemote(r))
	if err != nil {
		log.Fatal(err)
	}

	// pid => 127.0.0.1:4000/store/scraper
	e.Spawn(newStore, "store", actor.WithID("scraper"))

	select {}
}
