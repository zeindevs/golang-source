package router

import (
	"go-crud-redis/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(router *fiber.App, novelControleler *controller.NovelController) *fiber.App {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello bro")
	})

	router.Post("/novel", novelControleler.CreateNovel)
	router.Get("/novel/:id", novelControleler.GetNovelById)

	return router
}
