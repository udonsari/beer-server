package commands

import (
	"log"

	"github.com/urfave/cli/v2"

	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	userRepo "github.com/UdonSari/beer-server/domain/user/repo"
	"github.com/UdonSari/beer-server/main/server"
	_ "github.com/go-sql-driver/mysql"
)

type migrationUpCommand struct {
	d *server.Dependency
}

func NewMigrationUpCommand(d *server.Dependency) *migrationUpCommand {
	return &migrationUpCommand{
		d: d,
	}
}

func (c *migrationUpCommand) Command() *cli.Command {
	return &cli.Command{
		Name:   "migrate-up",
		Usage:  "Used to migrate up table used by server",
		Action: c.main,
	}
}

func (c *migrationUpCommand) main(ctx *cli.Context) error {
	result := c.d.MysqlDB().AutoMigrate(&userRepo.DBUser{}, &beerRepo.DBBeer{}, &beerRepo.DBComment{}, &beerRepo.DBRate{})
	if result.Error != nil {
		log.Printf("failed migration up %+v", result.Error)
	} else {
		log.Printf("succeeded migration up")
	}
	return nil
}
