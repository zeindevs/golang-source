package main

import (
	"net/http"

	"github.com/zeindevs/launch"
	"github.com/zeindevs/launch/examples/handler"
)

func main() {
	l := launch.New()
	l.Plug(RequestIdPlugin, ErrorPlug)
	l.Post("/user", handler.CreateUser)
	l.Start()
}

func RequestIdPlugin(next launch.Handler) launch.Handler {
	return func(c *launch.Ctx) error {
		c.Request().Header.Add("X-REQUEST-ID", "test")
		return next(c)
	}
}

func ErrorPlug(next launch.Handler) launch.Handler {
	return func(c *launch.Ctx) error {
		c.ErrorHandler = func(c *launch.Ctx, err error) {
			c.Status(http.StatusFailedDependency).JSON(map[string]string{"error": err.Error()})
		}
		return next(c)
	}
}
