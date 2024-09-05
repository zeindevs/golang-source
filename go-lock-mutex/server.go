package server

import (
	"fmt"
)

type Player struct {
	Name string
}

type SetFooMsg struct {
	value int
}

type GameState struct {
	players []*Player
	foo     int

	msgch chan any
}

func (g *GameState) Receive(msg any) {
	g.msgch <- msg
}

func (g *GameState) loop() {
	for msg := range g.msgch {
		g.handleMessage(msg)
	}
}

func (g *GameState) handleMessage(msg any) {
	switch m := msg.(type) {
	case *Player:
		g.addPlayer(m)
	case *SetFooMsg:
		g.handleSetFoo(m)
	default:
		panic("invalid message received")
	}
}

func (g *GameState) handleSetFoo(foo *SetFooMsg) {
	g.foo = foo.value

	fmt.Println("setting foo:", foo.value)
}

func (g *GameState) addPlayer(player *Player) {
	g.players = append(g.players, player)

	fmt.Println("adding player:", player.Name)
}

func NewGameState() *GameState {
	g := &GameState{
		players: []*Player{},
		msgch:   make(chan any, 10),
	}

	go g.loop()

	return g
}

type Server struct {
	gameState *GameState
}

func NewServer() *Server {
	return &Server{
		gameState: NewGameState(),
	}
}

func (s *Server) handleNewPlayer(player *Player) error {
	s.gameState.Receive(player)
	return nil
}

func (s *Server) handleSetFoo(val int) error {
	s.gameState.Receive(&SetFooMsg{value: val})
	return nil
}
