package handler

import (
	"github.com/zeindevs/launch"
	"github.com/zeindevs/launch/examples/data"
	"github.com/zeindevs/launch/examples/handler/params"
)

func CreateUser(c *launch.Ctx) error {
	params, err := launch.RequestParams[params.CreateUser](c)
	if err != nil {
		return err
	}
	user := data.CreateUser(params)
	return c.JSON(user)
}
