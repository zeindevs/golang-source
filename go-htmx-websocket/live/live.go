package live

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	updatePeriod = 1 * time.Second
	pingPeriod   = 10 * time.Second
	writeWait    = 10 * time.Second
)

type Client struct {
	Conn         *websocket.Conn
	Notification *Notification
	Send         chan []byte
}

func (c *Client) Pump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.Notification.Unregister <- c
		c.Conn.Close()
	}()
	for {
		select {
		case msg := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

type Notification struct {
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
}

func NewNotification() *Notification {
	return &Notification{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    map[*Client]bool{},
	}
}

func (n *Notification) Run() {
	for {
		select {
		case msg := <-n.Broadcast:
			for client := range n.Clients {
				select {
				case client.Send <- msg:
				default:
					close(client.Send)
					delete(n.Clients, client)
				}
			}
		case client := <-n.Register:
			n.Clients[client] = true
		case client := <-n.Unregister:
			if _, ok := n.Clients[client]; ok {
				delete(n.Clients, client)
				close(client.Send)
			}
		}
	}
}
