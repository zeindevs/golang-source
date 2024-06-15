package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func NewRouter() *gin.Engine {
  router := gin.Default()

  // Add swagger
  router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  router.GET("", func(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, "Hello Bro")
  })

  return router
}
