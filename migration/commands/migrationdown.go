package commands

import (
	"log"

	"github.com/urfave/cli/v2"

	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	userRepo "github.com/UdonSari/beer-server/domain/user/repo"
	"github.com/UdonSari/beer-server/main/server"
	_ "github.com/go-sql-driver/mysql"
)

type migrationDownCommand struct {
	d *server.Dependency
}

func NewMigrationDownCommand(d *server.Dependency) *migrationDownCommand {
	return &migrationDownCommand{
		d: d,
	}
}

func (c *migrationDownCommand) Command() *cli.Command {
	return &cli.Command{
		Name:   "migrate-down",
		Usage:  "Used to migrate down table used by server",
		Action: c.main,
	}
}

func (c *migrationDownCommand) main(ctx *cli.Context) error {
	result := c.d.MysqlDB(true).DropTableIfExists(&userRepo.DBUser{}, &beerRepo.DBBeer{}, &beerRepo.DBReview{})
	if result.Error != nil {
		log.Printf("failed migration down %+v", result.Error)
	} else {
		log.Printf("succeeded migration down")
	}
	return nil
}
