package main

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//go:embed ui/build/client
var staticFS embed.FS

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(filesystem.New(filesystem.Config{
		Root:       http.FS(staticFS),
		PathPrefix: "ui/build/client",
	}))
	app.Get("*", func(c *fiber.Ctx) error {
		return filesystem.SendFile(c, http.FS(staticFS), "ui/build/client/index.html")
	})
	app.Listen(":9001")
}
