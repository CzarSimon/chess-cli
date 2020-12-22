package engine

import (
	"fmt"

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
