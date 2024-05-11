package router

import (
	"ginjwt/controller"
	"ginjwt/middleware"
	"ginjwt/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserRepository           repository.UsersRepository
	AuthenticationController *controller.AuthenticationController
	UserController           *controller.UserController
}

func NewRouter(rc *RouterConfig) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome Home")
	})

	router := service.Group("/api")

	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", rc.AuthenticationController.Register)
	authenticationRouter.POST("/login", rc.AuthenticationController.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("", middleware.DeserializeUser(rc.UserRepository), rc.UserController.GetUsers)

	return service
}
