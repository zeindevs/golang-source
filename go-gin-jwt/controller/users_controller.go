package controller

import (
	"ginjwt/data/response"
	"ginjwt/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepository repository.UsersRepository
}

func NewUserController(repository repository.UsersRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users := c.userRepository.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Successfully fetch all user data!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
