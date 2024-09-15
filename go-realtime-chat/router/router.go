package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zeindevs/go-realtime-chat/internal/user"
	"github.com/zeindevs/go-realtime-chat/internal/ws"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

  r.LoadHTMLGlob("www/*")
  r.Static("/static", "www")

  r.GET("/", func(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "index.html", nil)
  })

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
  r.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
