package main

import (
	"log"

	"github.com/UdonSari/beer-server/main/server"

	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	userRepo "github.com/UdonSari/beer-server/domain/user/repo"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	d := server.NewDependency()
	result := d.MysqlDB().AutoMigrate(&userRepo.DBUser{}, &beerRepo.DBBeer{}, &beerRepo.DBComment{}, &beerRepo.DBRate{})
	if result.Error != nil {
		log.Printf("failed migration %+v", result.Error)
	} else {
		log.Printf("succeeded migration")
	}
}
