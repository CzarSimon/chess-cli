package cli

import (
	"fmt"
	"time"

	"github.com/CzarSimon/chess-cli/internal/engine"
	"github.com/CzarSimon/chess-cli/pkg/clio"
	"github.com/notnil/chess"
	"github.com/urfave/cli/v2"
)

const (
	displayModeDraw    = "draw"
	displayModeCommand = "command"
)

var (
	gameDisplayMode = ""
)

func createGameCommand(app *App) *cli.Command {
	return &cli.Command{
		Name:   "game",
		Usage:  "Creates a new game",
		Action: app.createGame,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "display-mode",
				Aliases:     []string{"m"},
				Usage:       fmt.Sprintf("Method to display the state of the game. Options: %s, %s", displayModeDraw, displayModeCommand),
				Value:       displayModeDraw,
				Destination: &gameDisplayMode,
			},
		},
	}
}

func (a *App) createGame(c *cli.Context) error {
	game := a.gameSvc.NewGame()
	engine := &engine.MinimaxEngine{Depth: 2}
	startGame(game, engine)

	return nil
}

func startGame(game *chess.Game, engine engine.Interface) {
	fmt.Println("Started game")
	if gameDisplayMode == displayModeDraw {
		draw(game)
	}

	for {
		ended := runPly(game, engine)
		if ended {
			break
		}
	}

	summarizeGame(game)
}

func runPly(game *chess.Game, engine engine.Interface) bool {
	if whiteToMove(game) {
		getMove(game)
	} else {
		generateMove(game, engine)
	}

	return game.Outcome() != chess.NoOutcome
}

func getMove(game *chess.Game) {
	input := clio.MustGet(explainTurn(game))

	err := game.MoveStr(input)
	if err != nil {
		fmt.Printf("Invalid move %s please try again\n\n", input)
		getMove(game)
	}
}

func generateMove(game *chess.Game, engine engine.Interface) {
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
	displayGeneratedMove(move, game)
}

func displayGeneratedMove(move string, game *chess.Game) {
	if gameDisplayMode == displayModeDraw {
		draw(game)
		return
	}

	fmt.Printf("Blacks move: %s\n\n", move)
}

func draw(game *chess.Game) {
	fmt.Println(game.Position().Board().Draw())
}

func summarizeGame(game *chess.Game) {
	fmt.Println("\nGame ended")
	fmt.Println("----------")
	outcome := "Draw"
	if game.Outcome() == chess.WhiteWon {
		outcome = "White won"
	} else if game.Outcome() == chess.BlackWon {
		outcome = "Black won"
	}

	fmt.Printf("%s by %s\n\n", outcome, game.Method())

	fmt.Printf("Moves:")
	fmt.Println(game)
}

func whiteToMove(game *chess.Game) bool {
	return game.Position().Turn() == chess.White
}

func explainTurn(game *chess.Game) string {
	return fmt.Sprintf("%s to move", game.Position().Turn().Name())
}
