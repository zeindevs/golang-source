package main

import (
	"github.com/zeindevs/launch"
	"github.com/zeindevs/launch/examples/handler"
)

func main() {
	launch.Post("/post", handler.CreatePost)
	launch.Get("/post/:id", handler.GetPost)
	launch.Start()
}
