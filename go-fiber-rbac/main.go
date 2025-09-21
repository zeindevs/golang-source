package main

import (
	"go-fiber-rbac/config"
	"go-fiber-rbac/db"
	"go-fiber-rbac/middleware"
	"go-fiber-rbac/models"
	"go-fiber-rbac/types"
	"go-fiber-rbac/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	cfg := config.New(log)

	db := db.Connect(cfg, log)

	app := fiber.New()
	app.Use(logger.New())

	app.Post("/auth/login", func(c *fiber.Ctx) error {
		var req types.LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}
		var user models.User
		if err := db.First(&user, "username = ?", req.Username).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid username or password",
			})
		}

		if err := util.ComparePassword(req.Password, user.Password); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid username or password",
			})
		}

		token, err := util.GenerateJWT(cfg, jwt.MapClaims{
			"uid": user.ID,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error",
			})
		}

		return c.JSON(fiber.Map{
			"data": token,
		})
	})

	app.Use(middleware.JWT(cfg))

	app.Get("/posts", middleware.Authorize(db, types.ViewPost), func(c *fiber.Ctx) error {
		return c.SendString("Viewing posts")
	})
	app.Post("/posts", middleware.MinLevel(db, types.LevelEditor), func(c *fiber.Ctx) error {
		return c.SendString("Post created")
	})
	app.Delete("/users/:id", middleware.MinLevel(db, types.LevelAdmin), func(c *fiber.Ctx) error {
		return c.SendString("User deleted")
	})

	app.Listen(cfg.GetString("LISTEN_ADDR"))
}
