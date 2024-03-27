package main

import (
	"log/slog"

	"github.com/gorilla/websocket"
)

type Peer interface {
	Send([]byte) error
}

type WSPeer struct {
	conn   *websocket.Conn
	server *Server
}

func NewWSPeer(conn *websocket.Conn, s *Server) *WSPeer {
	p := &WSPeer{
		conn:   conn,
		server: s,
	}
	go p.readLoop()
	return p

}

// FIXME: watch out for the memory leak later on!
func (p *WSPeer) readLoop() {
	var msg WSMessage
	for {
		if err := p.conn.ReadJSON(&msg); err != nil {
			slog.Error("ws peer read error", "err", err)
			return
		}
		if err := p.handleMessage(msg); err != nil {
			slog.Error("ws peer handle msg error", "err", err)
			return
		}
	}
}

func (p *WSPeer) handleMessage(msg WSMessage) error {
	if msg.Action == "subscribe" {
		p.server.AddPeerToTopics(p, msg.Topics...)
	}
	return nil
}

func (p *WSPeer) Send(b []byte) error {
	return p.conn.WriteMessage(websocket.BinaryMessage, b)
}
