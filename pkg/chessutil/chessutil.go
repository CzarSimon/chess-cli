package chessutil

import "github.com/notnil/chess"

// EncodeMove describes a move with algebraic notation
func EncodeMove(game *chess.Game, move *chess.Move) string {
	notation := chess.AlgebraicNotation{}
	return notation.Encode(game.Position(), move)
}
