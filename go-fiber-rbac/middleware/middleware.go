package middleware

import (
	"go-fiber-rbac/models"
	"go-fiber-rbac/types"
	"go-fiber-rbac/util"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Authorize(db *gorm.DB, required types.Permission) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("uid").(uint)

		var user models.User
		err := db.Preload("Role.Permissions").First(&user, userID).Error
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		for _, perm := range user.Role.Permissions {
			if perm.Name == string(required) {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "forbidden: insufficient permissions",
		})
	}
}

func MinLevel(db *gorm.DB, required types.Level) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("uid").(uint)

		var user models.User
		err := db.Preload("Role").First(&user, userID).Error
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "user not found",
			})
		}

		if user.Role.Level >= int(required) {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "forbidden: insufficient access level",
		})
	}
}

func JWT(cfg *viper.Viper) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("authorization")
		if authorization == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "token required",
			})
		}

		tokens := strings.Split(authorization, " ")
		if len(tokens) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid token",
			})
		}

		claims, err := util.ParseJWT(cfg, tokens[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		userID := uint(claims["uid"].(float64))
		c.Locals("uid", userID)

		return c.Next()
	}
}
