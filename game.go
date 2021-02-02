package kroki

import "math/rand"

const (
	DefaultNumTeeth = 10
)

type Game struct {
	Teeth    []bool `json:"teeth"`
	Lost     bool   `json:"lost"`
	BadTooth int    `json:"-"`
}

type Message struct {
	Event   string      `json:"e,omitempty"`
	Payload interface{} `json:"p,omitempty"`
}

func NewGame() *Game {
	return &Game{
		Teeth:    make([]bool, DefaultNumTeeth),
		BadTooth: rand.Intn(DefaultNumTeeth - 1),
	}
}
