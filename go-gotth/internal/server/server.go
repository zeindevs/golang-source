package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-gotth/internal/store"
	"github.com/go-gotth/internal/templates"
)

type GuestStore interface {
	AddGuest(guest store.Guest) error
	GetGuests() ([]store.Guest, error)
}

type Server struct {
	logger     *log.Logger
	port       string
	httpServer *http.Server
	guestDb    *store.GuestStore
}

func NewServer(logger *log.Logger, port string, guestDb *store.GuestStore) (*Server, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is required")
	}
	if guestDb == nil {
		return nil, fmt.Errorf("guestDb is required")
	}
	return &Server{
		logger:  logger,
		port:    port,
		guestDb: guestDb,
	}, nil
}

func (s *Server) Start() error {
	s.logger.Printf("starting server on port %s", s.port)

	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	router.HandleFunc("GET /", s.defaultHandler)

	s.httpServer = &http.Server{
		Addr:    s.port,
		Handler: router,
	}

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatalf("error when running server: %s", err)
		}
	}()

	<-stopCh

	if err := s.httpServer.Shutdown(context.Background()); err != nil {
		s.logger.Fatalf("error when shutting down server: %v", err)
		return err
	}

	return nil
}

func (s *Server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	templates.Index().Render(r.Context(), w)
}
