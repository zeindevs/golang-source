package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"sync"
)

type Message struct {
	data []byte
	from string
}

type Server struct {
	mu        sync.Mutex
	peers     map[net.Conn]bool
	addPeerCh chan net.Conn
	quitch    chan struct{}
	ln        net.Listener
	msgCh     chan Message
}

func NewServer() *Server {
	return &Server{
		peers:     make(map[net.Conn]bool),
		addPeerCh: make(chan net.Conn),
		quitch:    make(chan struct{}),
		msgCh:     make(chan Message),
	}
}

func (s *Server) loop() {
	for {
		select {
		case msg := <-s.msgCh:
			// we cant be here
			if err := s.handleMessage(msg); err != nil {
				fmt.Println(err)
			}
		case conn := <-s.addPeerCh:
			// we cant be here
			slog.Info("new peer connected", "add", conn.RemoteAddr())
			s.peers[conn] = true
			go s.readLoop(conn)
		case <-s.quitch:
			// if we are here
			for peer := range s.peers {
				peer.Close()
			}
			fmt.Println("gracefull server shutdown yeah!!!")
			s.ln.Close()
			return
		}
	}
}

func (s *Server) readLoop(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read loop error", err)
			continue
		}
		s.msgCh <- Message{
			from: conn.RemoteAddr().String(),
			data: buf[:n],
		}
	}
}

func (s *Server) handleMessage(msg Message) error {
	fmt.Println("got msg -> ", string(msg.data))
	for peer := range s.peers {
		if peer.RemoteAddr().String() != msg.from {
			peer.Write(msg.data)
		}
	}
	return nil
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		return err
	}
	s.ln = ln

	go s.loop()

	slog.Info("tcp server started", "port", ":3000")
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			fmt.Println("err =>", err.Error())
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	s.addPeerCh <- conn

	// code
	fmt.Println("handling connection:", conn.RemoteAddr())
}

func (s *Server) Close() error {
	s.quitch <- struct{}{}
	return nil
}

func main() {
	s := NewServer()

	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	s.Close()
	// }()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
