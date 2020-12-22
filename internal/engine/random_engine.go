package engine

import (
	"math/rand"
	"time"

	"github.com/CzarSimon/chess-cli/pkg/chessutil"
	"github.com/notnil/chess"
)

// RandomEngine implementation of a chess engine that plays random moves.
type RandomEngine struct {
	Seed   int64
	seeded bool
}

// NextMove randomly generates the next move
func (e *RandomEngine) NextMove(game *chess.Game) (string, error) {
	e.seed()

	moves := game.ValidMoves()
	if len(moves) < 1 {
		return "", ErrNoValidMove
	}

	moveIdx := rand.Intn(len(moves))
	return chessutil.EncodeMove(game, moves[moveIdx]), nil
}

func (e *RandomEngine) seed() {
	if e.seeded {
		return
	}

	seed := e.Seed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	rand.Seed(seed)
}
