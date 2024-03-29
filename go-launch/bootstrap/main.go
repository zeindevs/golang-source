package main

import (
	"github.com/zeindevs/launch"
	"github.com/zeindevs/launch/bootstrap/handlers"
)

func main() {
	// isAPI := os.Getenv("LAUNCH_API_PROJECT")
	app := launch.New(false)
	app.Get("/user", handlers.GetUserIndex)
	app.Get("/foo", func(c *launch.Ctx) error { return c.JSON("non JSON string") })
	// listenAddr := os.Getenv("LAUNCH_APP_ADDR")
	app.Start(":3000")
}

func MyCustomNotFoundPlug(next launch.Handler) launch.Handler {
	return func(c *launch.Ctx) error {
		return next(c)
	}
}
