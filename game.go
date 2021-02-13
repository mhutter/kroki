package kroki

import (
	"log"
	"math/rand"
)

const (
	// NumTeeth is the number of teeth allocated in new games
	NumTeeth int = 10
)

// Game represents an individual session, or a "Kroki"
type Game struct {
	ID string `json:"id"`

	Teeth    []bool `json:"teeth"`
	BadTooth int    `json:"-"`

	Players []*Player `json:"players"`
	WhoLost string    `json:"lost,omitempty"`
}

// NewGame returns a new, initialized Game
func NewGame() *Game {
	return &Game{
		Teeth:    make([]bool, NumTeeth),
		BadTooth: rand.Intn(NumTeeth),
		Players:  make([]*Player, 0),
	}
}

func (g *Game) Leave(p *Player) {
	for i := 0; i < len(g.Players); i++ {
		if g.Players[i].connID != p.connID {
			continue
		}
		// Slice out player
		g.Players = append(g.Players[:i], g.Players[i+1:]...)
		return

	}
}

func (g *Game) Join(p *Player) {
	// TODO: remove duplicate players in case a player has the app open in two
	// different browser windows
	for i, v := range g.Players {
		if v.ID == p.ID {
			// Replace existing player with same ID
			g.Players[i] = p
			// Disconnect old connection
			v.conn.WriteJSON(&Message{Event: "leave"})
			v.conn.Close()
			return
		}
	}
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
