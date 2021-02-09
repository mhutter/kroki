package kroki

import (
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
	Event   string      `json:"e,omitempty"`
	Payload interface{} `json:"p,omitempty"`
}

func (s Server) readLoop(conn *websocket.Conn) {
	s.numClients.Inc()
	defer s.numClients.Dec()

	var game *Game
	player := Player{conn: conn}

	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Println("Error reading message:", err)
			}
			return
		}

		log.Printf("message: %#v", msg)

		switch msg.Event {
		case "setName":
			player.Name = msg.Payload.(string)
		case "setPlayerID":
			player.ID = msg.Payload.(string)

		case "joinGame":
			game = s.handleJoinGame(msg)
			game.Join(&player)
			defer game.RemovePlayer(player.ID)
			game.Broadcast()

		case "press":
			id := int(msg.Payload.(float64))
			game.Teeth[id] = true
			if id == game.BadTooth {
				game.WhoLost = player.ID
			}
			game.Broadcast()

		case "restart":
			game.Restart()
			game.Broadcast()

		default:
			log.Printf("unknown event: %v", msg)
		}
	}
}

func (s Server) handleJoinGame(msg Message) *Game {
	var gid string

	if msg.Payload != nil {
		gid = msg.Payload.(string)
	}

	game, err := s.games.GetOrCreate(gid)
	if err != nil {
		log.Panicln(err)
	}

	return game
}
