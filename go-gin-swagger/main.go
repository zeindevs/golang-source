package main

import (
	_ "ginswagger/docs"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// @title       Tag Service API
// @version     1.0
// @description A Tag service API in Go using Gin framework

// @host localhost:8888
// @BasePath /api
func main() {
	log.Info().Msg("Started Server!")
	s := gin.New()

	s.Run(":8888")
}
