package cli

import (
	"fmt"
	"time"

	"github.com/CzarSimon/chess-cli/internal/engine"
	"github.com/CzarSimon/chess-cli/pkg/clio"
	"github.com/notnil/chess"
	"github.com/urfave/cli/v2"
)

type turn bool

func (t turn) String() string {
	if t == white {
		return "White"
	}

	return "Black"
}

func (t turn) Int() int {
	if t == white {
		return 0
	}

	return 1
}

func (t turn) Explain() string {
	return fmt.Sprintf("%s to move", t)
}

func (t turn) Next() turn {
	return !t
}

const (
	white turn = false
	black turn = true
)

func createGameCommand(app *App) *cli.Command {
	return &cli.Command{
		Name:   "game",
		Usage:  "Creates a new game",
		Action: app.createGame,
	}
}

func (a *App) createGame(c *cli.Context) error {
	game := a.gameSvc.NewGame()
	engine := &engine.RandomEngine{}
	startGame(game, engine)

	return nil
}

func startGame(game *chess.Game, engine engine.Interface) {
	fmt.Println("Started game")
	currentTurn := white

	for {
		ended := runPly(currentTurn, game, engine)
		if ended {
			break
		}

		currentTurn = currentTurn.Next()
	}
}

func runPly(t turn, game *chess.Game, engine engine.Interface) bool {
	if t == white {
		draw(game)
		getMove(t, game)
	} else {
		generateMove(t, game, engine)
	}

	return game.Outcome() != chess.NoOutcome
}

func getMove(t turn, game *chess.Game) {
	input := clio.MustGet(t.Explain())
	err := game.MoveStr(input)
	if err != nil {
		fmt.Printf("Invalid move %s. Reason: %v\n", input, err)
		getMove(t, game)
	}
}

func generateMove(t turn, game *chess.Game, engine engine.Interface) {
	move, err := engine.NextMove(game)
	if err != nil {
		panic(err)
	}

	err = game.MoveStr(move)
	if err != nil {
		fmt.Printf("Invalid move %s. Reason: %v\n", move, err)
		panic(err)
	}
	time.Sleep(200 * time.Millisecond)
}

func draw(game *chess.Game) {
	fmt.Println(game.Position().Board().Draw())
}
