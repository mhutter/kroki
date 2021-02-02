package kroki

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	handler  http.Handler
	upgrader *websocket.Upgrader
	games    GameRepo

	// Metrics
	numClients prometheus.Gauge
}

type GameRepo interface {
	Create() (string, *Game)
	Get(string) (*Game, error)
}

func NewServer() Server {
	var (
		upgrader   = &websocket.Upgrader{}
		mux        = http.NewServeMux()
		games      = NewRepo()
		numClients = promauto.NewGauge(prometheus.GaugeOpts{
			Namespace: "kroki",
			Subsystem: "websocket",
			Name:      "num_clients",
			Help:      "Number of currently connected clients",
		})
		s = Server{mux, upgrader, games, numClients}
	)

	upgrader.CheckOrigin = func(_ *http.Request) bool { return true }

	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/_healthz", s.handleHealthz)
	mux.HandleFunc("/ws", s.handleWS)
	mux.Handle("/", http.FileServer(http.Dir("public")))

	return s
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

func (s Server) handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (s Server) handleWS(w http.ResponseWriter, r *http.Request) {

	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	s.readLoop(conn)
}

func (s Server) readLoop(conn *websocket.Conn) {
	s.numClients.Inc()
	defer s.numClients.Dec()
	g := NewGame()
	sendSetate(conn, g)

	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			if !websocket.IsCloseError(err, websocket.CloseGoingAway) {
				log.Println("Error reading message:", err)
			}
			return
		}

		switch msg.Event {
		case "press":
			id := int(msg.Payload.(float64))
			g.Teeth[id] = true
			if id == g.BadTooth {
				g.Lost = true
			}
		}

		sendSetate(conn, g)
	}
}

func sendSetate(conn *websocket.Conn, game *Game) {
	msg := &Message{Event: "update", Payload: game}
	if err := conn.WriteJSON(msg); err != nil {
		log.Println("Error writing response:", err)
	}
}
