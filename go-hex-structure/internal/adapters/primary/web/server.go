package web

import (
	"fmt"
	"hex-structure/internal/core/services/users"
	"log"
	"strconv"
)

type App struct {
  UserService *users.Service
  ListenAddr string
}

func NewApp(userService *users.Service, port string) *App {
  return &App{
    UserService: userService,
    ListenAddr: port,
  }
}

func WithPort(port int) string {
  return fmt.Sprintf(":%s", strconv.Itoa(port))
}

func (a *App) Run() {
  log.Println("Web server listening on port:", a.ListenAddr)
}
