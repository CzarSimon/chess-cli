package cli

import (
	"fmt"
	"os"

	"github.com/CzarSimon/chess-cli/internal/service"
	"github.com/urfave/cli/v2"
)

const (
	appName    = "chess"
	appVersion = "<version>" // Do not manually change this line. Its updated automatically by scripts/build-cli.sh
)

// App cli application
type App struct {
	gameSvc *service.GameService
}

// New creates a cli application instance.
func New() *App {
	return &App{
		gameSvc: &service.GameService{},
	}
}

// Run runs the cli application and logs error if one occured
func (a *App) Run() {
	err := a.createCLI().Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func (a *App) createCLI() *cli.App {
	return &cli.App{
		Name:    appName,
		Version: appVersion,
		Usage:   "Cli chess game with a built in engine",
		Commands: []*cli.Command{
			{
				Name: "create",
				Subcommands: []*cli.Command{
					createGameCommand(a),
				},
			},
		},
	}
}
