package main

import (
	"log"
	"log/slog"
	"net/http"
	"strings"
)

type Producer interface {
	Start() error
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HTTPProducer struct {
	listenAddr string
	producech  chan<- Message
}

func NewHTTPProducer(listenAddr string, producech chan Message) *HTTPProducer {
	return &HTTPProducer{
		listenAddr: listenAddr,
		producech:  producech,
	}
}

func (p *HTTPProducer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		path  = strings.TrimPrefix(r.URL.Path, "/")
		parts = strings.Split(path, "/")
	)
	// commit
	if r.Method == "GET" {

	}

	// publishing
	if r.Method == "POST" {
		if len(parts) != 2 {
			log.Printf("invalid action")
			return
		}
		topic := parts[1]
		p.producech <- Message{
			Topic: topic,
			Data:  []byte("We don't know yet"),
		}
	}

	log.Println(parts)
	// w.Write([]byte(r.URL.Path))
}

func (p *HTTPProducer) Start() error {
	slog.Info("HTTP transport started", "port", p.listenAddr)
	return http.ListenAndServe(p.listenAddr, p)
}
