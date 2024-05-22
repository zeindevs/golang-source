package users

import (
	"hex-structure/internal/core/services/users"

	"github.com/gofiber/fiber"
)

func (s *Service) CreateAccount(c *fiber.Ctx) error {
	ctx := c.Context()

	var input struct {
		Username string `json:"username"`
	}
	if err := c.BodyParser(&input); err != nil {
		c.SendStatus(400)
	}

	resp, err := s.userAPI.CreateAccount(ctx, users.CreateAccountReq(input))
  if err != nil {
    c.SendStatus(500)
  }

  var out struct {
    UserID string `json:"userId"`
  }
  out.UserID = resp.UserID
  return c.JSON(out)
}
