package api

import "net/http"

func (s *Server) handleFooA(w http.ResponseWriter, r *http.Request) {}
func (s *Server) handleFooB(w http.ResponseWriter, r *http.Request) {}
func (s *Server) handleFooC(w http.ResponseWriter, r *http.Request) {}
