package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	addr   string
	router *http.ServeMux
}

type ServerConfig struct {
	Addr   string
	Router *http.ServeMux
}

func NewServer(cfg *ServerConfig) *Server {
	s := &Server{
		addr:   cfg.Addr,
		router: cfg.Router,
	}
	return s
}

func (s *Server) Run() error {
	slog.Info("server up and listening on", "url", fmt.Sprintf("http://localhost%s", s.addr))
	return http.ListenAndServe(s.addr, s.router)
}
