package main

import (
	"github.com/zeindevs/go-consumer-interface-patterns/internal/bar"
	"github.com/zeindevs/go-consumer-interface-patterns/internal/foo"
)

func main() {
	bar := &bar.Bar{}
	foo := foo.NewFoo(bar)

	foo.Greet()
}
