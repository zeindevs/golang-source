package internal

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	egress     chan Event
	chatroom   string
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	c.connection.SetReadLimit(512)

	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}

	c.connection.SetPongHandler(c.pongHandler)

	for {
		messageType, payload, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}
		log.Println("MessageType:", messageType)
		log.Println("Payload:", string(payload))

		var reqEvent Event
		if err := json.Unmarshal(payload, &reqEvent); err != nil {
			log.Printf("error marshalling message: %v", err)
			break
		}

		if err := c.manager.routeEvent(reqEvent, c); err != nil {
			log.Println("error handeling Message:", err.Error())
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	log.Println("pong", pongMsg)
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

func (c *Client) writeMessages() {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		c.manager.removeClient(c)
	}()

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed:", err.Error())
				}

				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err.Error())
				return
			}

			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err.Error())
			}
			log.Println("send message")
		case <-ticker.C:
			log.Println("ping")
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("writemsg:", err.Error())
				return
			}
		}
	}
}
