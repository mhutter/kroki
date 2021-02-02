package kroki

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var (
	idCharset = []rune(`abcdefghijklmnopqrstuvwxyz0123456789`)
)

// Repo stores and returns Games
type Repo struct {
	games map[string]*Game
	rand  *rand.Rand
}

var _ GameRepo = &Repo{}

// NewRepo returns a new, empty Repo with its RNG (for ID generation)
// initialized.
func NewRepo() *Repo {
	return &Repo{
		games: make(map[string]*Game),
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Create creates a new game with a random ID.
func (r *Repo) Create() (id string, game *Game) {
	game = NewGame()
	id = r.newGameID()
	r.games[id] = game
	return
}

// Get retrieves the game with the given ID, or an error if the given game does
// not exist.
func (r *Repo) Get(id string) (*Game, error) {
	if g := r.games[id]; g != nil {
		return g, nil
	}

	return nil, fmt.Errorf("Game '%s' not found", id)
}

// newGameID returns a random Game ID that is not in use yet.
func (r Repo) newGameID() string {
	for {
		id := randomString(r.rand, 4, idCharset)
		if r.games[id] == nil {
			return id
		}
	}
}

func randomString(rng *rand.Rand, size int, charset []rune) string {
	var buf strings.Builder
	for i := 0; i < size; i++ {
		n := rng.Intn(len(idCharset))
		buf.WriteRune(idCharset[n])
	}
	return buf.String()
}
