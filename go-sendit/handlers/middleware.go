package handlers

import "github.com/gofiber/fiber/v2"

func AuthMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func MustAuthMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SubdomainCheckerMiddleware(c *fiber.Ctx) error {
	return c.Next()
}
