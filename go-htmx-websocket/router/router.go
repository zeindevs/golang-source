package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/zeindevs/go-htmx-websocket/handler"
)

func NewRouter(h *handler.Handler, staticDir http.Dir) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(staticDir)))
	r.Get("/", h.DashboardHandler)
	r.Get("/users", h.UserHandler)
	r.Get("/users/new", h.NewUsersHandler)
	r.Post("/users", h.CreateUserHandler)
	r.Get("/users/{id}", h.UserHandler)
	r.Get("/users-table", h.UsersTableHandler)
	r.Get("/refresh-chart/{id}", h.RefreshChartHandler)
	r.Get("/live", h.Live)
	r.Post("/notification", h.Notification)
	return r
}
