package engine

import (
	"fmt"

	"github.com/notnil/chess"
)

var pieceValues map[chess.PieceType]float32 = map[chess.PieceType]float32{
	chess.Pawn:   1,
	chess.Knight: 3,
	chess.Bishop: 3,
	chess.Rook:   5,
	chess.Queen:  9,
	chess.King:   900,
}

// ScoredMove represents a move and the score that has been assigned to it.
type ScoredMove struct {
	Move  *chess.Move
	Score float32
}

func (m ScoredMove) String() string {
	return fmt.Sprintf("ScoredMove(move=%s, score=%.2f)", m.Move, m.Score)
}

// Score returns a score evaluating a chess possition
func Score(pos *chess.Position) float32 {
	board := pos.Board().SquareMap()
	var score float32

	for square, piece := range board {
		score += scorePiece(piece, square)
	}

	return score
}

func scorePiece(piece chess.Piece, square chess.Square) float32 {
	var coeff float32 = 1
	if piece.Color() == chess.Black {
		coeff = -1
	}

	t := piece.Type()
	pieceValue, _ := pieceValues[t]

	return coeff * pieceValue
}
