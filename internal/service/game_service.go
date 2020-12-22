package service

import "github.com/notnil/chess"

// GameService service responsible for game business logic
type GameService struct{}

// NewGame creates a new chess game.
func (gs *GameService) NewGame() *chess.Game {
	return chess.NewGame()
}
