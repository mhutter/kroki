package kroki

import (
	"github.com/gorilla/websocket"
)

type Player struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	conn *websocket.Conn
}
