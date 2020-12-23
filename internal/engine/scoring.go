package engine

import (
	"fmt"

	"github.com/notnil/chess"
)

type squareValues map[chess.Square]float32

var pieceValues map[chess.PieceType]float32 = map[chess.PieceType]float32{
	chess.Pawn:   10,
	chess.Knight: 30,
	chess.Bishop: 30,
	chess.Rook:   50,
	chess.Queen:  90,
	chess.King:   900,
}

var whitePawnValues = squareValues{
	chess.A8: 0.0, chess.B8: 0.0, chess.C8: 0.0, chess.D8: 0.0, chess.E8: 0.0, chess.F8: 0.0, chess.G8: 0.0, chess.H8: 0.0,
	chess.A7: 5.0, chess.B7: 5.0, chess.C7: 5.0, chess.D7: 5.0, chess.E7: 5.0, chess.F7: 5.0, chess.G7: 5.0, chess.H7: 5.0,
	chess.A6: 1.0, chess.B6: 1.0, chess.C6: 2.0, chess.D6: 3.0, chess.E6: 3.0, chess.F6: 2.0, chess.G6: 1.0, chess.H6: 1.0,
	chess.A5: 0.5, chess.B5: 0.5, chess.C5: 1.0, chess.D5: 2.5, chess.E5: 2.5, chess.F5: 1.0, chess.G5: 0.5, chess.H5: 0.5,
	chess.A4: 0.0, chess.B4: 0.0, chess.C4: 0.0, chess.D4: 2.0, chess.E4: 2.0, chess.F4: 0.0, chess.G4: 0.0, chess.H4: 0.0,
	chess.A3: 0.5, chess.B3: -0.5, chess.C3: -1.0, chess.D3: 0.0, chess.E3: 0.0, chess.F3: -1.0, chess.G3: -0.5, chess.H3: 0.5,
	chess.A2: 0.5, chess.B2: 1.0, chess.C2: 1.0, chess.D2: -2.0, chess.E2: -2.0, chess.F2: 1.0, chess.G2: 1.0, chess.H2: 0.5,
	chess.A1: 0.0, chess.B1: 0.0, chess.C1: 0.0, chess.D1: 0.0, chess.E1: 0.0, chess.F1: 0.0, chess.G1: 0.0, chess.H1: 0.0,
}

var blackPawnValues = squareValues{
	chess.A8: 0.0, chess.B8: 0.0, chess.C8: 0.0, chess.D8: 0.0, chess.E8: 0.0, chess.F8: 0.0, chess.G8: 0.0, chess.H8: 0.0,
	chess.A7: 0.5, chess.B7: 1.0, chess.C7: 1.0, chess.D7: -2.0, chess.E7: -2.0, chess.F7: 1.0, chess.G7: 1.0, chess.H7: 5.0,
	chess.A6: 0.5, chess.B6: -0.5, chess.C6: -1.0, chess.D6: 0.0, chess.E6: 0.0, chess.F6: -1.0, chess.G6: -0.5, chess.H6: 0.5,
	chess.A5: 0.0, chess.B5: 0.0, chess.C5: 0.0, chess.D5: 2.0, chess.E5: 2.0, chess.F5: 0.0, chess.G5: 0.0, chess.H5: 0.0,
	chess.A4: 0.5, chess.B4: 0.5, chess.C4: 1.0, chess.D4: 2.5, chess.E4: 2.5, chess.F4: 1.0, chess.G4: 0.5, chess.H4: 0.5,
	chess.A3: 1.0, chess.B3: 1.0, chess.C3: 2.0, chess.D3: 3.0, chess.E3: 3.0, chess.F3: 2.0, chess.G3: 1.0, chess.H3: 1.0,
	chess.A2: 5.0, chess.B2: 5.0, chess.C2: 5.0, chess.D2: 5.0, chess.E2: 5.0, chess.F2: 5.0, chess.G2: 5.0, chess.H2: 5.0,
	chess.A1: 0.0, chess.B1: 0.0, chess.C1: 0.0, chess.D1: 0.0, chess.E1: 0.0, chess.F1: 0.0, chess.G1: 0.0, chess.H1: 0.0,
}

var whiteKnightValues = squareValues{
	chess.A8: -5.0, chess.B8: -4.0, chess.C8: -3.0, chess.D8: -3.0, chess.E8: -3.0, chess.F8: -3.0, chess.G8: -4.0, chess.H8: -5.0,
	chess.A7: -4.0, chess.B7: -2.0, chess.C7: 0.0, chess.D7: 0.0, chess.E7: 0.0, chess.F7: 0.0, chess.G7: -2.0, chess.H7: -4.0,
	chess.A6: -3.0, chess.B6: 0.0, chess.C6: 1.0, chess.D6: 1.5, chess.E6: 1.5, chess.F6: 1.0, chess.G6: 0.0, chess.H6: -3.0,
	chess.A5: -3.0, chess.B5: 0.5, chess.C5: 1.5, chess.D5: 2.0, chess.E5: 2.0, chess.F5: 1.5, chess.G5: 0.5, chess.H5: -3.0,
	chess.A4: -3.0, chess.B4: 0.0, chess.C4: 1.5, chess.D4: 2.0, chess.E4: 2.0, chess.F4: 1.0, chess.G4: 0.0, chess.H4: -3.0,
	chess.A3: -3.0, chess.B3: 0.5, chess.C3: 1.0, chess.D3: 1.5, chess.E3: 1.5, chess.F3: 1.0, chess.G3: 0.5, chess.H3: -3.0,
	chess.A2: -4.0, chess.B2: -2.0, chess.C2: 0.0, chess.D2: 0.5, chess.E2: 0.5, chess.F2: 0.0, chess.G2: -2.0, chess.H2: -4.0,
	chess.A1: -5.0, chess.B1: -4.0, chess.C1: -3.0, chess.D1: -3.0, chess.E1: -3.0, chess.F1: -3.0, chess.G1: -4.0, chess.H1: -5.0,
}

var blackKnightValues = squareValues{
	chess.A8: -5.0, chess.B8: -4.0, chess.C8: -3.0, chess.D8: -3.0, chess.E8: -3.0, chess.F8: -3.0, chess.G8: -4.0, chess.H8: -5.0,
	chess.A7: -4.0, chess.B7: -2.0, chess.C7: 0.0, chess.D7: 0.5, chess.E7: 0.5, chess.F7: 0.0, chess.G7: -2.0, chess.H7: -4.0,
	chess.A6: -3.0, chess.B6: 0.5, chess.C6: 1.0, chess.D6: 1.5, chess.E6: 1.5, chess.F6: 1.0, chess.G6: 0.5, chess.H6: -3.0,
	chess.A5: -3.0, chess.B5: 0.0, chess.C5: 1.5, chess.D5: 2.0, chess.E5: 2.0, chess.F5: 1.0, chess.G5: 0.0, chess.H5: -3.0,
	chess.A4: -3.0, chess.B4: 0.5, chess.C4: 1.5, chess.D4: 2.0, chess.E4: 2.0, chess.F4: 1.5, chess.G4: 0.5, chess.H4: -3.0,
	chess.A3: -3.0, chess.B3: 0.0, chess.C3: 1.0, chess.D3: 1.5, chess.E3: 1.5, chess.F3: 1.0, chess.G3: 0.0, chess.H3: -3.0,
	chess.A2: -4.0, chess.B2: -2.0, chess.C2: 0.0, chess.D2: 0.0, chess.E2: 0.0, chess.F2: 0.0, chess.G2: -2.0, chess.H2: -4.0,
	chess.A1: -5.0, chess.B1: -4.0, chess.C1: -3.0, chess.D1: -3.0, chess.E1: -3.0, chess.F1: -3.0, chess.G1: -4.0, chess.H1: -5.0,
}

var whiteBishopValues = squareValues{
	chess.A8: -2.0, chess.B8: -1.0, chess.C8: -1.0, chess.D8: -1.0, chess.E8: -1.0, chess.F8: -1.0, chess.G8: -1.0, chess.H8: -2.0,
	chess.A7: -1.0, chess.B7: 0.0, chess.C7: 0.0, chess.D7: 0.0, chess.E7: 0.0, chess.F7: 0.0, chess.G7: 0.0, chess.H7: -1.0,
	chess.A6: -1.0, chess.B6: 0.0, chess.C6: 0.5, chess.D6: 1.0, chess.E6: 1.0, chess.F6: 0.5, chess.G6: 0.0, chess.H6: -1.0,
	chess.A5: -1.0, chess.B5: 0.5, chess.C5: 0.5, chess.D5: 1.0, chess.E5: 1.0, chess.F5: 0.5, chess.G5: 0.5, chess.H5: -1.0,
	chess.A4: -1.0, chess.B4: 0.0, chess.C4: 1.0, chess.D4: 1.0, chess.E4: 1.0, chess.F4: 1.0, chess.G4: 0.0, chess.H4: -1.0,
	chess.A3: -1.0, chess.B3: 1.0, chess.C3: 1.0, chess.D3: 1.0, chess.E3: 1.0, chess.F3: 1.0, chess.G3: 1.0, chess.H3: -1.0,
	chess.A2: -1.0, chess.B2: 0.5, chess.C2: 0.0, chess.D2: 0.0, chess.E2: 0.0, chess.F2: 0.0, chess.G2: 0.5, chess.H2: -1.0,
	chess.A1: -2.0, chess.B1: -1.0, chess.C1: -1.0, chess.D1: -1.0, chess.E1: -1.0, chess.F1: -1.0, chess.G1: -1.0, chess.H1: -2.0,
}

var blackBishopValues = squareValues{
	chess.A8: -2.0, chess.B8: -1.0, chess.C8: -1.0, chess.D8: -1.0, chess.E8: -1.0, chess.F8: -1.0, chess.G8: -1.0, chess.H8: -2.0,
	chess.A7: -1.0, chess.B7: 0.5, chess.C7: 0.0, chess.D7: 0.0, chess.E7: 0.0, chess.F7: 0.0, chess.G7: 0.5, chess.H7: -1.0,
	chess.A6: -1.0, chess.B6: 1.0, chess.C6: 1.0, chess.D6: 1.0, chess.E6: 1.0, chess.F6: 1.0, chess.G6: 1.0, chess.H6: -1.0,
	chess.A5: -1.0, chess.B5: 0.0, chess.C5: 1.0, chess.D5: 1.0, chess.E5: 1.0, chess.F5: 1.0, chess.G5: 0.0, chess.H5: -1.0,
	chess.A4: -1.0, chess.B4: 0.5, chess.C4: 0.5, chess.D4: 1.0, chess.E4: 1.0, chess.F4: 0.5, chess.G4: 0.5, chess.H4: -1.0,
	chess.A3: -1.0, chess.B3: 0.0, chess.C3: 0.5, chess.D3: 1.0, chess.E3: 1.0, chess.F3: 0.5, chess.G3: 0.0, chess.H3: -1.0,
	chess.A2: -1.0, chess.B2: 0.0, chess.C2: 0.0, chess.D2: 0.0, chess.E2: 0.0, chess.F2: 0.0, chess.G2: 0.0, chess.H2: -1.0,
	chess.A1: -2.0, chess.B1: -1.0, chess.C1: -1.0, chess.D1: -1.0, chess.E1: -1.0, chess.F1: -1.0, chess.G1: -1.0, chess.H1: -2.0,
}

var whiteRookValues = squareValues{
	chess.A8: 0.0, chess.B8: 0.0, chess.C8: 0.0, chess.D8: 0.0, chess.E8: 0.0, chess.F8: 0.0, chess.G8: 0.0, chess.H8: 0.0,
	chess.A7: 0.5, chess.B7: 1.0, chess.C7: 1.0, chess.D7: 1.0, chess.E7: 1.0, chess.F7: 1.0, chess.G7: 1.0, chess.H7: 0.5,
	chess.A6: -0.5, chess.B6: 0.0, chess.C6: 0.0, chess.D6: 0.0, chess.E6: 0.0, chess.F6: 0.0, chess.G6: 0.0, chess.H6: -0.5,
	chess.A5: -0.5, chess.B5: 0.0, chess.C5: 0.0, chess.D5: 0.0, chess.E5: 0.0, chess.F5: 0.0, chess.G5: 0.0, chess.H5: -0.5,
	chess.A4: -0.5, chess.B4: 0.0, chess.C4: 0.0, chess.D4: 0.0, chess.E4: 0.0, chess.F4: 0.0, chess.G4: 0.0, chess.H4: -0.5,
	chess.A3: -0.5, chess.B3: 0.0, chess.C3: 0.0, chess.D3: 0.0, chess.E3: 0.0, chess.F3: 0.0, chess.G3: 0.0, chess.H3: -0.5,
	chess.A2: -0.5, chess.B2: 0.0, chess.C2: 0.0, chess.D2: 0.0, chess.E2: 0.0, chess.F2: 0.0, chess.G2: 0.0, chess.H2: -0.5,
	chess.A1: 0.0, chess.B1: 0.0, chess.C1: 0.0, chess.D1: 0.5, chess.E1: 0.5, chess.F1: 0.0, chess.G1: 0.0, chess.H1: 0.0,
}

var blackRookValues = squareValues{
	chess.A8: 0.0, chess.B8: 0.0, chess.C8: 0.0, chess.D8: 5.0, chess.E8: 5.0, chess.F8: 0.0, chess.G8: 0.0, chess.H8: 0.0,
	chess.A7: -0.5, chess.B7: 0.0, chess.C7: 0.0, chess.D7: 0.0, chess.E7: 0.0, chess.F7: 0.0, chess.G7: 0.0, chess.H7: -0.5,
	chess.A6: -0.5, chess.B6: 0.0, chess.C6: 0.0, chess.D6: 0.0, chess.E6: 0.0, chess.F6: 0.0, chess.G6: 0.0, chess.H6: -0.5,
	chess.A5: -0.5, chess.B5: 0.0, chess.C5: 0.0, chess.D5: 0.0, chess.E5: 0.0, chess.F5: 0.0, chess.G5: 0.0, chess.H5: -0.5,
	chess.A4: -0.5, chess.B4: 0.0, chess.C4: 0.0, chess.D4: 0.0, chess.E4: 0.0, chess.F4: 0.0, chess.G4: 0.0, chess.H4: -0.5,
	chess.A3: -0.5, chess.B3: 0.0, chess.C3: 0.0, chess.D3: 0.0, chess.E3: 0.0, chess.F3: 0.0, chess.G3: 0.0, chess.H3: -0.5,
	chess.A2: 0.5, chess.B2: 1.0, chess.C2: 1.0, chess.D2: 1.0, chess.E2: 1.0, chess.F2: 1.0, chess.G2: 1.0, chess.H2: 0.5,
	chess.A1: 0.0, chess.B1: 0.0, chess.C1: 0.0, chess.D1: 0.0, chess.E1: 0.0, chess.F1: 0.0, chess.G1: 0.0, chess.H1: 0.0,
}

var whiteQueenValues = squareValues{
	chess.A8: -2.0, chess.B8: -1.0, chess.C8: -1.0, chess.D8: -0.5, chess.E8: -0.5, chess.F8: -1.0, chess.G8: -1.0, chess.H8: -2.0,
	chess.A7: -1.0, chess.B7: 0.0, chess.C7: 0.0, chess.D7: 0.0, chess.E7: 0.0, chess.F7: 0.0, chess.G7: 0.0, chess.H7: -1.0,
	chess.A6: -1.0, chess.B6: 0.0, chess.C6: 0.5, chess.D6: 0.5, chess.E6: 0.5, chess.F6: 0.5, chess.G6: 0.0, chess.H6: -1.0,
	chess.A5: -0.5, chess.B5: 0.0, chess.C5: 0.5, chess.D5: 0.5, chess.E5: 0.5, chess.F5: 0.5, chess.G5: 0.0, chess.H5: -0.5,
	chess.A4: 0.0, chess.B4: 0.0, chess.C4: 0.5, chess.D4: 0.5, chess.E4: 0.5, chess.F4: 0.5, chess.G4: 0.0, chess.H4: -0.5,
	chess.A3: -1.0, chess.B3: 0.0, chess.C3: 0.5, chess.D3: 0.5, chess.E3: 0.5, chess.F3: 0.5, chess.G3: 0.0, chess.H3: -1.0,
	chess.A2: -1.0, chess.B2: 0.0, chess.C2: 0.5, chess.D2: 0.0, chess.E2: 0.0, chess.F2: 0.0, chess.G2: 0.0, chess.H2: -1.0,
	chess.A1: -2.0, chess.B1: -1.0, chess.C1: -1.0, chess.D1: -0.5, chess.E1: -0.5, chess.F1: -1.0, chess.G1: -1.0, chess.H1: -2.0,
}

var blackQueenValues = squareValues{
	chess.A8: -2.0, chess.B8: -1.0, chess.C8: -1.0, chess.D8: -0.5, chess.E8: -0.5, chess.F8: -1.0, chess.G8: -1.0, chess.H8: -2.0,
	chess.A7: -1.0, chess.B7: 0.0, chess.C7: 0.0, chess.D7: 0.0, chess.E7: 0.0, chess.F7: 0.5, chess.G7: 0.0, chess.H7: -1.0,
	chess.A6: -1.0, chess.B6: 0.0, chess.C6: 0.5, chess.D6: 0.5, chess.E6: 0.5, chess.F6: 0.5, chess.G6: 0.0, chess.H6: -1.0,
	chess.A5: -0.5, chess.B5: 0.0, chess.C5: 0.5, chess.D5: 0.5, chess.E5: 0.5, chess.F5: 0.5, chess.G5: 0.0, chess.H5: 0.0,
	chess.A4: -0.5, chess.B4: 0.0, chess.C4: 0.5, chess.D4: 0.5, chess.E4: 0.5, chess.F4: 0.5, chess.G4: 0.0, chess.H4: -0.5,
	chess.A3: -1.0, chess.B3: 0.0, chess.C3: 0.5, chess.D3: 0.5, chess.E3: 0.5, chess.F3: 0.5, chess.G3: 0.0, chess.H3: -1.0,
	chess.A2: -1.0, chess.B2: 0.0, chess.C2: 0.0, chess.D2: 0.0, chess.E2: 0.0, chess.F2: 0.0, chess.G2: 0.0, chess.H2: -1.0,
	chess.A1: -2.0, chess.B1: -1.0, chess.C1: -1.0, chess.D1: -0.5, chess.E1: -0.5, chess.F1: -1.0, chess.G1: -1.0, chess.H1: -2.0,
}

var whiteKingValues = squareValues{
	chess.A8: 0.0, chess.B8: 0.0, chess.C8: 0.0, chess.D8: 0.0, chess.E8: 0.0, chess.F8: 0.0, chess.G8: 0.0, chess.H8: 0.0,
	chess.A7: 0.0, chess.B7: 0.0, chess.C7: 0.0, chess.D7: 0.0, chess.E7: 0.0, chess.F7: 0.0, chess.G7: 0.0, chess.H7: 0.0,
	chess.A6: 0.0, chess.B6: 0.0, chess.C6: 0.0, chess.D6: 0.0, chess.E6: 0.0, chess.F6: 0.0, chess.G6: 0.0, chess.H6: 0.0,
	chess.A5: 0.0, chess.B5: 0.0, chess.C5: 0.0, chess.D5: 0.0, chess.E5: 0.0, chess.F5: 0.0, chess.G5: 0.0, chess.H5: 0.0,
	chess.A4: 0.0, chess.B4: 0.0, chess.C4: 0.0, chess.D4: 0.0, chess.E4: 0.0, chess.F4: 0.0, chess.G4: 0.0, chess.H4: 0.0,
	chess.A3: 0.0, chess.B3: 0.0, chess.C3: 0.0, chess.D3: 0.0, chess.E3: 0.0, chess.F3: 0.0, chess.G3: 0.0, chess.H3: 0.0,
	chess.A2: 0.0, chess.B2: 0.0, chess.C2: 0.0, chess.D2: 0.0, chess.E2: 0.0, chess.F2: 0.0, chess.G2: 0.0, chess.H2: 0.0,
	chess.A1: 0.0, chess.B1: 0.0, chess.C1: 0.0, chess.D1: 0.0, chess.E1: 0.0, chess.F1: 0.0, chess.G1: 0.0, chess.H1: 0.0,
}

var blackKingValues = squareValues{
	chess.A8: 0.0, chess.B8: 0.0, chess.C8: 0.0, chess.D8: 0.0, chess.E8: 0.0, chess.F8: 0.0, chess.G8: 0.0, chess.H8: 0.0,
	chess.A7: 0.0, chess.B7: 0.0, chess.C7: 0.0, chess.D7: 0.0, chess.E7: 0.0, chess.F7: 0.0, chess.G7: 0.0, chess.H7: 0.0,
	chess.A6: 0.0, chess.B6: 0.0, chess.C6: 0.0, chess.D6: 0.0, chess.E6: 0.0, chess.F6: 0.0, chess.G6: 0.0, chess.H6: 0.0,
	chess.A5: 0.0, chess.B5: 0.0, chess.C5: 0.0, chess.D5: 0.0, chess.E5: 0.0, chess.F5: 0.0, chess.G5: 0.0, chess.H5: 0.0,
	chess.A4: 0.0, chess.B4: 0.0, chess.C4: 0.0, chess.D4: 0.0, chess.E4: 0.0, chess.F4: 0.0, chess.G4: 0.0, chess.H4: 0.0,
	chess.A3: 0.0, chess.B3: 0.0, chess.C3: 0.0, chess.D3: 0.0, chess.E3: 0.0, chess.F3: 0.0, chess.G3: 0.0, chess.H3: 0.0,
	chess.A2: 0.0, chess.B2: 0.0, chess.C2: 0.0, chess.D2: 0.0, chess.E2: 0.0, chess.F2: 0.0, chess.G2: 0.0, chess.H2: 0.0,
	chess.A1: 0.0, chess.B1: 0.0, chess.C1: 0.0, chess.D1: 0.0, chess.E1: 0.0, chess.F1: 0.0, chess.G1: 0.0, chess.H1: 0.0,
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
	color := piece.Color()
	if piece.Color() == chess.Black {
		coeff = -1
	}

	var value float32

	t := piece.Type()
	switch t {
	case chess.Pawn:
		value = calculatePawnValue(square, color)
	case chess.Knight:
		value = calculateKnightValue(square, color)
	case chess.Bishop:
		value = calculateBishopValue(square, color)
	case chess.Rook:
		value = calculateRookValue(square, color)
	case chess.Queen:
		value = calculateQueenValue(square, color)
	case chess.King:
		value = calculateKingValue(square, color)
	default:
		value = calculateDefaultValue(t)
	}

	return coeff * value
}

func calculatePawnValue(square chess.Square, color chess.Color) float32 {
	pieceValue, _ := pieceValues[chess.Pawn]

	if color == chess.Black {
		squareValue, _ := blackPawnValues[square]
		return pieceValue + squareValue
	}

	squareValue, _ := whitePawnValues[square]
	return pieceValue + squareValue
}

func calculateKnightValue(square chess.Square, color chess.Color) float32 {
	pieceValue, _ := pieceValues[chess.Knight]

	if color == chess.Black {
		squareValue, _ := blackKnightValues[square]
		return pieceValue + squareValue
	}

	squareValue, _ := whiteKnightValues[square]
	return pieceValue + squareValue
}

func calculateBishopValue(square chess.Square, color chess.Color) float32 {
	pieceValue, _ := pieceValues[chess.Knight]

	if color == chess.Black {
		squareValue, _ := blackBishopValues[square]
		return pieceValue + squareValue
	}

	squareValue, _ := whiteBishopValues[square]
	return pieceValue + squareValue
}

func calculateRookValue(square chess.Square, color chess.Color) float32 {
	pieceValue, _ := pieceValues[chess.Knight]

	if color == chess.Black {
		squareValue, _ := blackRookValues[square]
		return pieceValue + squareValue
	}

	squareValue, _ := whiteRookValues[square]
	return pieceValue + squareValue
}

func calculateQueenValue(square chess.Square, color chess.Color) float32 {
	pieceValue, _ := pieceValues[chess.Knight]

	if color == chess.Black {
		squareValue, _ := blackQueenValues[square]
		return pieceValue + squareValue
	}

	squareValue, _ := whiteQueenValues[square]
	return pieceValue + squareValue
}

func calculateKingValue(square chess.Square, color chess.Color) float32 {
	pieceValue, _ := pieceValues[chess.Knight]

	if color == chess.Black {
		squareValue, _ := blackKingValues[square]
		return pieceValue + squareValue
	}

	squareValue, _ := whiteKingValues[square]
	return pieceValue + squareValue
}

func calculateDefaultValue(t chess.PieceType) float32 {
	value, _ := pieceValues[t]
	return value
}
