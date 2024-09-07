package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Username string
	Role     string
}

func main() {
	app := fiber.New()

	app.Get("/post", handleGetPost)                                  // public
	app.Post("/post/manage", onlyAdmin(handleGetPostManage))         // admin
	app.Post("/post/special", onlySpecialUser(handleGetPostSpecial)) // special

	log.Fatal(app.Listen(":4000"))
}

func onlyAdmin(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := getUserFromDB()
		if user.Role != "admin" {
			return c.SendStatus(http.StatusUnauthorized)
		}
		return fn(c)
	}
}

func onlySpecialUser(fn fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := getUserFromDB()
		if user.Role == "admin" || user.Role == "special" {
			return fn(c)
		}
		return c.SendStatus(http.StatusUnauthorized)
	}
}

func getUserFromDB() User {
	return User{
		Username: "James",
		Role:     "admin",
	}
}

func handleGetPost(c *fiber.Ctx) error {
	return c.JSON("some posts here")
}

func handleGetPostManage(c *fiber.Ctx) error {
	return c.JSON("the admin page of this post")
}

func handleGetPostSpecial(c *fiber.Ctx) error {
	return c.JSON("the special page of this post")
}
