package main

import (
	"fmt"
	"os"

	"github.com/UdonSari/beer-server/main/server"
	"github.com/UdonSari/beer-server/migration/commands"
	"github.com/urfave/cli/v2"
)

func getCommands(d *server.Dependency) []*cli.Command {
	return []*cli.Command{
		commands.NewMigrationUpCommand(d).Command(),
		commands.NewMigrationDownCommand(d).Command(),
		commands.NewSeedFakeCommand(d).Command(),
		commands.NewSeedCommand(d).Command(),
		commands.NewSeedManualCommand(d).Command(),
	}
}

func main() {
	d := server.NewDependency()

	app := &cli.App{
		Name:     "command executor",
		Commands: getCommands(&d),
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
