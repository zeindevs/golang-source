package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Service struct {
	store Storage
}

func NewService(store Storage) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users/{userID}/parse-kindle-file", s.handleParseKindleFile).Methods(http.MethodPost)
	router.HandleFunc("/cloud/send-daily-insights", s.handleSendDailyInsights).Methods(http.MethodGet)
}

func (s *Service) handleParseKindleFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	file, _, err := r.FormFile("file")
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing file: %v", err))
		return
	}

	defer file.Close()

	// parse that multipart file
	raw, err := parseKindleExtractFile(file)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing file: %v", err))
		return
	}

	id, _ := strconv.Atoi(userID)

	if err := s.createDataFromRawBook(raw, id); err != nil {
		WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error creating data from raw book: %v", err))
		return
	}

	WriteJSON(w, http.StatusCreated, map[string]any{"message": "Successfully parsed file"})
}

func (s *Service) handleSendDailyInsights(w http.ResponseWriter, r *http.Request) {
	// get users

	// loop over users and get random highlights
	// build an email and sent it
}

func (s *Service) createDataFromRawBook(raw *RawExtractBook, userID int) error {
	_, err := s.store.GetBookByISBN(raw.ASIN)
	if err != nil {
		s.store.CreateBook(Book{
			ISBN:    raw.ASIN,
			Title:   raw.Title,
			Authors: raw.Authors,
		})
	}

	// create highlights
	hs := make([]Highlight, len(raw.Highlights))
	for i, h := range raw.Highlights {
		hs[i] = Highlight{
			Text:     h.Text,
			Location: h.Location.URL,
			Note:     h.Note,
			UserID:   userID,
			BookID:   raw.ASIN,
		}
	}

	if err = s.store.CreateHighlights(hs); err != nil {
		log.Printf("Error creating highlights: %v\n", err)
		return err
	}

	return nil
}

func parseKindleExtractFile(file multipart.File) (*RawExtractBook, error) {
	deocoder := json.NewDecoder(file)

	raw := new(RawExtractBook)
	if err := deocoder.Decode(raw); err != nil {
		return nil, err
	}

	return raw, nil
}
