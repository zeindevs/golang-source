package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zeindevs/go-project-structure/db"
	"github.com/zeindevs/go-project-structure/handlers"
	"github.com/zeindevs/go-project-structure/pkg/saltedge"
	"github.com/zeindevs/go-project-structure/types"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.HTTPErrorHandler = handlers.HTTPErrorHandler

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}
	if err := saltedge.Init(); err != nil {
		log.Fatal(err)
	}

	e.Use(handlers.WithLoggin)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	api := e.Group("/api", handlers.WithAuthentication)
	api.GET("/user", handlers.HandleGetUser)
	api.POST("/onboard", handlers.HandlePostOnboard)
	api.GET("/auction", handlers.HandleGetAuctions)

	sellside := api.Group("", handlers.WithAccountTypeOnly(types.AccountTypeSellSide))
	sellside.GET("/metrics/sellside", handlers.HandleSellSideDemoMetrics)
	sellside.POST("/funding", handlers.HandlePostFundingRequest)
	sellside.GET("/funding", handlers.HandleGetFundingRequest)

	e.Start(os.Getenv("LISTEN_ADDR"))
}
