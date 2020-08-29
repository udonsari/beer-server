package main

import (
	"log"
	"math/rand"
	"strconv"

	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	"github.com/UdonSari/beer-server/main/server"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	d := server.NewDependency()
	mysqlDB := d.MysqlDB()

	for i := 0; i < 100; i++ {
		dbBeer := getRandomBeer()
		log.Printf("trying to put %vth beer %v", i, spew.Sdump(dbBeer))
		if err := mysqlDB.Create(&dbBeer).Error; err != nil {
			log.Fatalf("failed to add %+v with err %+v", spew.Sdump(dbBeer), spew.Sdump(err))
		}
	}

	for i := 0; i < 500; i++ {
		dbComment := getRandomComment()
		log.Printf("trying to put %vth comment %v", i, spew.Sdump(dbComment))
		if err := mysqlDB.Create(&dbComment).Error; err != nil {
			log.Fatalf("failed to add %+v with err %+v", spew.Sdump(dbComment), spew.Sdump(err))
		}

		dbRate := getRandomRate()
		log.Printf("trying to put %vth rate %v", i, spew.Sdump(dbRate))
		if err := mysqlDB.Create(&dbRate).Error; err != nil {
			log.Fatalf("failed to add %+v with err %+v", spew.Sdump(dbRate), spew.Sdump(err))
		}
	}
}

func getRandomBeer() beerRepo.DBBeer {
	name := "TEST_NAME_" + strconv.Itoa(rand.Int())
	brewery := "TEST_BREWAERY_" + strconv.Itoa(rand.Int()%1000)
	abv := rand.Float64() * 10
	country := "TEST_COUNTRY_" + strconv.Itoa(rand.Int()%30)
	beerStyle := "TEST_STYLE_" + strconv.Itoa(rand.Int()%20)

	aromaList := ""
	for i := 0; i < 3; i++ {
		aromaList += "TEST_AROMA_" + strconv.Itoa(rand.Int()%10)
		if i != 2 {
			aromaList += "___"
		}
	}

	imageURLList := ""
	for i := 0; i < 5; i++ {
		imageURLList += "http:naver.com"
		if i != 4 {
			imageURLList += "___"
		}
	}

	return beerRepo.DBBeer{
		Name:         name,
		Brewery:      brewery,
		ABV:          abv,
		Country:      country,
		BeerStyle:    beerStyle,
		AromaList:    aromaList,
		ImageURLList: imageURLList,
	}
}

func getRandomComment() beerRepo.DBComment {
	return beerRepo.DBComment{
		BeerID:  rand.Int63n(100),
		Content: "TEST_COMMENT_" + strconv.Itoa(rand.Int()),
		UserID:  rand.Int63n(100),
	}
}

func getRandomRate() beerRepo.DBRate {
	return beerRepo.DBRate{
		BeerID: rand.Int63n(100),
		Ratio:  rand.Float64() * 5,
		UserID: rand.Int63n(100),
	}
}
