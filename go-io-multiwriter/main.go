package main

import (
	"bytes"
	"fmt"
	"io"
)

type Conn struct {
	io.Writer
}

func NewConn() *Conn {
	return &Conn{
		Writer: new(bytes.Buffer),
	}
}

func (c *Conn) Write(b []byte) (int, error) {
	fmt.Println("writing to underlying connection:", string(b))
	return c.Writer.Write(b)
}

type Server struct {
	peers map[*Conn]bool
}

func NewServer() *Server {
	s := &Server{
		peers: make(map[*Conn]bool),
	}

	for i := 0; i < 10; i++ {
		s.peers[NewConn()] = true
	}

	return s
}

func (s *Server) broadcast(msg []byte) error {
	peers := []io.Writer{}
	for peer := range s.peers {
		peers = append(peers, peer)
	}
	mw := io.MultiWriter(peers...)
	_, err := mw.Write(msg)
	return err
}

func main() {
	s := NewServer()
	s.broadcast([]byte("foo"))
}
