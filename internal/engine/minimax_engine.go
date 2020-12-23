package engine

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/CzarSimon/chess-cli/pkg/chessutil"
	"github.com/notnil/chess"
)

// MinimaxEngine implementation of a chess engine that uses
// minimax tree search to evaluate which next move is the best.
type MinimaxEngine struct {
	Depth  uint
	Seed   int64
	seeded bool
}

// NextMove generates the next move by recursively scoring possible moves with a given depth.
func (e *MinimaxEngine) NextMove(game *chess.Game) (string, error) {
	moves := game.ValidMoves()
	if len(moves) < 1 {
		return "", ErrNoValidMove
	}

	scoredMoves := scoreMoves(game.Position(), moves, e.Depth)
	bestMove := e.selectBestMove(scoredMoves)

	return chessutil.EncodeMove(game, bestMove), nil
}

func (e *MinimaxEngine) selectBestMove(moves []ScoredMove) *chess.Move {
	bestMoves := make([]ScoredMove, 0)
	var bestScore float32 = math.MaxFloat32

	for _, move := range moves {
		if move.Score == bestScore {
			bestMoves = append(bestMoves, move)
		} else if move.Score < bestScore {
			bestMoves = []ScoredMove{move}
			bestScore = move.Score
		}
	}

	e.seed()
	// displayBestMoves(bestScore, bestMoves)
	moveIdx := rand.Intn(len(bestMoves))
	return bestMoves[moveIdx].Move
}

func scoreMoves(pos *chess.Position, moves []*chess.Move, depth uint) []ScoredMove {
	if depth == 0 {
		return scorePositions(pos, moves)
	}

	scores := make([]ScoredMove, len(moves))
	for i, move := range moves {
		nextPos := pos.Update(move)
		nextMoves := pos.ValidMoves()
		if len(nextMoves) == 0 {
			scores[i] = ScoredMove{
				Move:  move,
				Score: Score(nextPos),
			}
			continue
		}

		nextScores := scoreMoves(nextPos, nextMoves, depth-1)
		scores[i] = selectHighestScoredMove(nextScores)
	}

	return scores
}

func selectHighestScoredMove(moves []ScoredMove) ScoredMove {
	bestMove := moves[0]
	bestScore := moves[0].Score

	for _, move := range moves {
		if move.Score > bestScore {
			bestMove = move
			bestScore = move.Score
		}
	}

	return bestMove
}

func scorePositions(pos *chess.Position, moves []*chess.Move) []ScoredMove {
	scores := make([]ScoredMove, len(moves))

	for i, move := range moves {
		next := pos.Update(move)
		score := ScoredMove{
			Move:  move,
			Score: Score(next),
		}
		scores[i] = score
	}

	return scores
}

func (e *MinimaxEngine) seed() {
	if e.seeded {
		return
	}

	seed := e.Seed
	if seed == 0 {
		seed = time.Now().UnixNano()
	}

	rand.Seed(seed)
}

func displayScore(move ScoredMove, pos *chess.Position) {
	moveStr := chess.AlgebraicNotation{}.Encode(pos, move.Move)
	fmt.Printf("move: %s. score: %f%s\n", moveStr, move.Score, pos.Board().Draw())
}

func displayBestMoves(score float32, moves []ScoredMove) {
	fmt.Printf("best score: %.2f. moves:\n", score)
	for i, m := range moves {
		fmt.Printf("%d. %s\n", i+1, m)
	}
}
