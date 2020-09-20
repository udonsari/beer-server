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
	result := c.d.MysqlDB(true).AutoMigrate(&userRepo.DBUser{}, &beerRepo.DBBeer{}, &beerRepo.DBReview{})
	if result.Error != nil {
		log.Printf("failed migration up %+v", result.Error)
	} else {
		log.Printf("succeeded migration up")
	}

	// TODO Model에서 `gorm:"uniqueIndex:beer_id_user_id,sort:desc"` 이 태그가 안먹는데, 더 알아보기
	// User is allowed to add review only one time for same beer
	result = c.d.MysqlDB(true).Model(&beerRepo.DBReview{}).AddUniqueIndex("beer_id_user_id", "beer_id", "user_id")
	if result.Error != nil {
		log.Printf("failed adding unique index %+v", result.Error)
	} else {
		log.Printf("succeeded adding unique index")
	}
	return nil
}
