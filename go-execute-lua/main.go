package main

import (
	"log"

	"github.com/Shopify/go-lua"
)

var script = `
bestBid = bestBid();
bestAsk = bestAsk();
spread = math.abs(bestBid - bestAsk);
print(spread);
`

func main() {
	l := lua.NewState()
	lua.OpenLibraries(l)
	registerBestAsk(l)
	registerBestBid(l)
	if err := lua.DoString(l, script); err != nil {
		log.Fatal(err)
	}
}

func registerBestBid(l *lua.State) {
	l.Register("bestBid", func(s *lua.State) int {
		l.PushInteger(10000)
		return 1
	})
}

func registerBestAsk(l *lua.State) {
	l.Register("bestAsk", func(s *lua.State) int {
		l.PushInteger(10001)
		return 1
	})
}
