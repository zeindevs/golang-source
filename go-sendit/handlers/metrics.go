package handlers

import "github.com/gofiber/fiber/v2"

func HandleMetrics(c *fiber.Ctx) error {
	return c.Next()
}
