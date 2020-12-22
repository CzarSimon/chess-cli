package engine

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/notnil/chess"
)

// Common errors
var (
	ErrNoValidMove = fmt.Errorf("no valid moves")
)

// Interface chess engine interface.
type Interface interface {
	NextMove(game *chess.Game) (string, error)
}

// RandomEngine implementation of a chess engine that plays random moves.
type RandomEngine struct {
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
	move := moves[moveIdx]

	return move.String(), nil
}

func (e *RandomEngine) seed() {
	if e.seeded {
		return
	}

	rand.Seed(time.Now().UnixNano())
}
