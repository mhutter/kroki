package kroki

import (
	"github.com/gorilla/websocket"
)

// Player represents an individual client interacting with a game.
type Player struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	conn *websocket.Conn
}
