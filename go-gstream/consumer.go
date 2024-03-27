package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

func foo() {
	websocket.DefaultDialer.Dial("ws://oo", nil)
}

var upgrader = websocket.Upgrader{}

type Consumer interface {
	Start() error
}

type WSConsumer struct {
	ListenAddr string
	server     *Server
}

func NewWSConsumer(listenAddr string, server *Server) *WSConsumer {
	return &WSConsumer{
		ListenAddr: listenAddr,
		server:     server,
	}
}

func (ws *WSConsumer) Start() error {
	slog.Info("websocket consumer started", "port", ws.ListenAddr)
	return http.ListenAndServe(ws.ListenAddr, ws)
}

func (ws *WSConsumer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	p := NewWSPeer(conn, ws.server)
	ws.server.AddConn(p)
	// fmt.Println(conn)
}

type WSMessage struct {
	Action string   `json:"action"`
	Topics []string `json:"topics"`
}
