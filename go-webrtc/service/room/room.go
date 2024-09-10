package room

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

func NewRoomMap() *RoomMap {
	return &RoomMap{
		Mutex: sync.RWMutex{},
		Map:   map[string][]Participant{},
	}
}

func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)

	r.Map[roomID] = []Participant{}

	return roomID
}

func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) error {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	_, ok := r.Map[roomID]
	if !ok {
		return fmt.Errorf("Room not found")
	}

	p := Participant{host, conn}

	log.Println("Inserting into roomID:", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)

	return nil
}

func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)
}

func (r *RoomMap) DeleteParticipant(roomID string, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	for i, p := range r.Map[roomID] {
		if p.Conn == conn {
			r.Map[roomID] = r.Map[roomID][:i]
			break
		}
	}
}
