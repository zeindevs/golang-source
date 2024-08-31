package main

import (
	"encoding/json"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/anthdm/hollywood/actor"
	"github.com/anthdm/hollywood/remote"
	"github.com/zeindevs/go-hollymicro/types"
)

const (
	scrapeInterval = time.Second
	url            = "https://catfact.ninja/fact"
)

type scraper struct {
	url      string
	storePID *actor.PID
	engine   *actor.Engine
}

func newScraper(url string, storePID *actor.PID) actor.Producer {
	return func() actor.Receiver {
		return &scraper{
			url:      url,
			storePID: storePID,
		}
	}
}

func (s *scraper) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		slog.Debug("scraper is started")
		s.engine = c.Engine()
		go s.scrapeLoop()
	case actor.Stopped:
		slog.Debug("scrape is stopped")
	default:
		slog.Warn("store got unknown message", "data", msg, "type", reflect.TypeOf(msg).String())
	}
}

type CatFact struct {
	Fact string `json:"fact"`
}

func (s *scraper) scrapeLoop() {
	for {
		resp, err := http.Get(s.url)
		if err != nil {
			panic(err)
		}
		var res CatFact
		if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
			slog.Error(err.Error())
			continue
		}
		s.engine.Send(s.storePID, &types.CatFact{
			Fact: res.Fact,
		})
		time.Sleep(scrapeInterval)
	}
}

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))

	listenAddr := flag.String("listenAddr", "127.0.0.1:3000", "todo")
	flag.Parse()

	r := remote.New(*listenAddr, remote.NewConfig())
	e, err := actor.NewEngine(actor.NewEngineConfig().WithRemote(r))
	if err != nil {
		log.Fatal(err)
	}

	// pid >> 127.0.0.1:4000/store/scraper
	storePID := actor.NewPID("127.0.0.1:4000", "store/scraper")

	e.Spawn(newScraper(url, storePID), "scraper")

	select {}
}
