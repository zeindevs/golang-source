package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
	"github.com/zeindevs/go-htmx-websocket/data"
	"github.com/zeindevs/go-htmx-websocket/live"
	"github.com/zeindevs/go-htmx-websocket/views"
)

type Handler struct {
	bundb        *bun.DB
	notification *live.Notification
}

func NewHandler(bundb *bun.DB, notification *live.Notification) *Handler {
	return &Handler{
		bundb:        bundb,
		notification: notification,
	}
}

func (h *Handler) UserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var user data.User

	if err := h.bundb.NewRaw("select id, name, email, status from users", bun.Ident("public_01_initial")); err != nil {
		log.Println("err get user:", err)
		return
	}
	log.Println(id)

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) NewUsersHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: NewUsersHandler
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: CreateUserHandler
}

func (h *Handler) UsersTableHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: UsersTableHandler
}

func (h *Handler) RefreshChartHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: RefreshChartHandler
}

func (h *Handler) DashboardHandler(w http.ResponseWriter, r *http.Request) {
	metrics := make([]data.Metric, 0)
	if err := data.GetMetric(r.Context(), h.bundb, &metrics); err != nil {
		log.Println(err)
		return
	}
	component := views.Dashboard(metrics)
	component.Render(r.Context(), w)
}

func (h *Handler) Live(w http.ResponseWriter, r *http.Request) {
	c, err := live.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("err upgrade:", err)
		return
	}
	client := &live.Client{
		Conn:         c,
		Notification: &live.Notification{},
		Send:         make(chan []byte),
	}
	h.notification.Register <- client
	go client.Pump()
}

func (h *Handler) Notification(w http.ResponseWriter, r *http.Request) {
	h.notification.Broadcast <- []byte(`<div id="notification" hx-swap-oob="beforeend"><p>New notification</p></div>`)
}
