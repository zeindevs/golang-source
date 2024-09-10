package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/zeindevs/go-webrtc/service/room"
	"github.com/zeindevs/go-webrtc/types"
	"github.com/zeindevs/go-webrtc/utils"
)

type broadcastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

type WSHandler struct {
	rooms     *room.RoomMap
	broadcast chan broadcastMsg
}

func NewWSHandler(rooms *room.RoomMap) *WSHandler {
	return &WSHandler{
		rooms:     rooms,
		broadcast: make(chan broadcastMsg),
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *WSHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("/create", s.handleCreateRoom)
	router.HandleFunc("/join", s.handleJoinRoom)
}

func (s *WSHandler) handleCreateRoom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	roomID := s.rooms.CreateRoom()

	resp := types.CreateRoomResponse{
		RoomID: roomID,
	}

	slog.Info("current rooms", "data", s.rooms.Map)

	utils.WriteJSON(w, http.StatusCreated, resp)
}

func (s *WSHandler) handleJoinRoom(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("roomID")
	isHost := true

	if roomID == "" {
		slog.Error("roomID missing in URL Parameters")
		return
	}

	if _, ok := s.rooms.Map[roomID]; !ok {
		slog.Error("room not found")
		return
	}

	if len(s.rooms.Map[roomID]) > 0 {
		isHost = false
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("web socket upgrade", "error", err)
		return
	}

	if err := s.rooms.InsertIntoRoom(roomID, isHost, ws); err != nil {
		ws.WriteJSON(map[string]any{"error": err.Error()})
		ws.Close()
		return
	}

	slog.Info("current rooms", "data", s.rooms.Map)

	quitCh := make(chan bool)

	go s.broadcaster(quitCh)

	for {
		var msg broadcastMsg

		if err := ws.ReadJSON(&msg.Message); err != nil {
			slog.Error("read ws data ", "error", err)
			s.rooms.DeleteParticipant(roomID, ws)
			quitCh <- true
			break
		}

		msg.Client = ws
		msg.RoomID = roomID

		s.broadcast <- msg
	}

	slog.Info("ws handler shutdown")
}

func (s *WSHandler) broadcaster(quitCh chan bool) {
loop:
	for {
		select {
		case msg := <-s.broadcast:
			for _, client := range s.rooms.Map[msg.RoomID] {
				if client.Conn != msg.Client {
					err := client.Conn.WriteJSON(msg.Message)
					if err != nil {
						slog.Error("send broadcaster data", "error", err)
						continue
					}
				}
			}
		case <-quitCh:
			break loop
		}
	}

	slog.Info("current rooms", "data", s.rooms.Map)
	slog.Info("broadcaster shutdown")
}
