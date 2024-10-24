package store

import (
	"encoding/json"
	"fmt"
	"log"
)

type Guest struct {
	Name  string
	Email string
}

type GuestStore struct {
	logger *log.Logger
	guests map[string]Guest
}

func NewGuestStore(logger *log.Logger) *GuestStore {
	return &GuestStore{
		logger: logger,
		guests: make(map[string]Guest),
	}
}

func (s *GuestStore) AddGuest(guest Guest) error {
	if guest.Email == "" {
		return fmt.Errorf("email is required")
	}
	if _, ok := s.guests[guest.Email]; ok {
		return fmt.Errorf("guest with email %s already exists", guest.Email)
	}
	s.guests[guest.Email] = guest
	fmt.Printf("added guest: %v \n", guest)
	return nil
}

func (s *GuestStore) GetGuests() ([]Guest, error) {
	if s.guests == nil {
		return nil, fmt.Errorf("no guests found")
	}
	guests := make([]Guest, 0, len(s.guests))
	for _, guest := range s.guests {
		guests = append(guests, guest)
	}
	return guests, nil
}

func DecodeGuest(payload []byte) (Guest, error) {
	var guest Guest
	if err := json.Unmarshal(payload, &guest); err != nil {
		return Guest{}, fmt.Errorf("error decoding guest: %w", err)
	}
	return guest, nil
}
