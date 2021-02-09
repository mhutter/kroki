package kroki

import (
	"log"
	"math/rand"
)

const (
	NumTeeth = 10
)

type Game struct {
	ID       string    `json:"id"`
	Teeth    []bool    `json:"teeth"`
	WhoLost  string    `json:"lost,omitempty"`
	Players  []*Player `json:"players"`
	BadTooth int       `json:"-"`
}

func NewGame() *Game {
	return &Game{
		Teeth:    make([]bool, NumTeeth),
		BadTooth: rand.Intn(NumTeeth - 1),
		Players:  make([]*Player, 0),
	}
}

func (g *Game) RemovePlayer(id string) {
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i].ID != id {
			continue
		}
		g.Players = append(g.Players[:i], g.Players[i+1:]...)
		return

	}
}

func (g *Game) Join(p *Player) {
	// TODO: remove duplicate players in case a player has the app open in two
	// different browser windows
	g.Players = append(g.Players, p)
}

func (g *Game) Restart() {
	g.Teeth = make([]bool, NumTeeth)
	g.BadTooth = rand.Intn(NumTeeth - 1)
	g.WhoLost = ""
}

func (g *Game) Broadcast() {
	msg := &Message{Event: "update", Payload: g}
	for _, p := range g.Players {
		if err := p.conn.WriteJSON(msg); err != nil {
			log.Println("Error writing response:", err)
		}
	}
}
