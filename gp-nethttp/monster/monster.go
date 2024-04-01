package monster

import (
	"encoding/json"
	"net/http"
)

type Monster struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Powers []string `json:"powers"`
}

type MonsterRequest struct {
	id     int
	name   string
	powers []string
}

var monsters []Monster

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		monster := Monster{
			ID:     1,
			Name:   "Drakula",
			Powers: []string{"Eat", "Kill", "Hug"},
		}

		monsters = append(monsters, monster)

		json.NewEncoder(w).Encode(monster)
	} else {
		json.NewEncoder(w).Encode(map[string]string{})
	}
}

func FindByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodGet {
		monster := Monster{}

		json.NewEncoder(w).Encode(monster)
	} else {
		json.NewEncoder(w).Encode(map[string]string{})
	}
}

func UpdateByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle get
	} else if r.Method == http.MethodPost {
		// Handle post
	} else if r.Method == http.MethodPut {
		// Handle put
	}
}

func DeleteByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Handle get
	} else if r.Method == http.MethodPost {
		// Handle post
	} else if r.Method == http.MethodPut {
		// Handle put
	}
}
