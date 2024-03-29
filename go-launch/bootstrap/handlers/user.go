package handlers

import (
	"github.com/zeindevs/launch"
	"github.com/zeindevs/launch/bootstrap/data"
	"github.com/zeindevs/launch/bootstrap/handlers/params"
)

func GetUserIndex(c *launch.Ctx) error {
	user := data.User{
		FirstName: "James",
		LastName:  "Jimmy",
		Email:     "james@jimmy.family",
		Roles:     []string{"user", "officer", "chad", "daddy"},
	}
	return c.View("errors/404", nil)
	return c.View("user/index", map[string]any{"user": user})
}

func CreateUser(c *launch.Ctx) error {
	params, err := launch.RequestParams[params.CreateUser](c)
	if err != nil {
		return err
	}
	user := data.CreateUser(params)
	return c.JSON(user)
}
