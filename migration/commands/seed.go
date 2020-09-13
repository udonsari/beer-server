package commands

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/UdonSari/beer-server/domain/beer"
	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	"github.com/UdonSari/beer-server/main/server"
	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
)

const (
	rateCommentNumber = 2000
	beerNumber        = 100
	rateBase          = 3
	abvLimit          = 10
	countryNumber     = 10
	beerStyleNumber   = 5
	breweryNumber     = 100
	aromaNumber       = 5
)

type seedCommand struct {
	d *server.Dependency
}

func NewSeedCommand(d *server.Dependency) *seedCommand {
	rand.Seed(time.Now().UnixNano())
	return &seedCommand{
		d: d,
	}
}

func (c *seedCommand) Command() *cli.Command {
	return &cli.Command{
		Name:   "seed",
		Usage:  "Used to seed data used by server",
		Action: c.main,
	}
}

func (c *seedCommand) main(ctx *cli.Context) error {
	br := beerRepo.New(c.d.MysqlDB(), 100)
	bu := beer.NewUseCase(br)

	for i := 0; i < beerNumber; i++ {
		beer := getRandomBeer()
		log.Printf("trying to put %vth beer %v", i, beer)
		if err := bu.Addbeer(beer); err != nil {
			log.Fatalf("failed to add %+v with err %+v", spew.Sdump(beer), spew.Sdump(err))
		}
	}

	for i := 0; i < 2000; i++ {
		comment := getRandomComment()
		log.Printf("trying to put %vth comment %v", i, comment)
		if err := bu.AddComment(comment); err != nil {
			log.Fatalf("failed to add %+v with err %+v", spew.Sdump(comment), spew.Sdump(err))
		}

		rate := getRandomRate()
		log.Printf("trying to put %vth rate %v", i, rate)
		if err := bu.AddRate(rate); err != nil {
			log.Fatalf("failed to add %+v with err %+v", spew.Sdump(rate), spew.Sdump(err))
		}
	}
	return nil
}

func getRandomBeer() beer.Beer {
	name := "TEST_NAME_" + strconv.Itoa(rand.Int())
	brewery := "TEST_BREWAERY_" + strconv.Itoa(rand.Int()%breweryNumber)
	abv := rand.Float64() * abvLimit
	country := "TEST_COUNTRY_" + strconv.Itoa(rand.Int()%countryNumber)
	beerStyle := "TEST_STYLE_" + strconv.Itoa(rand.Int()%beerStyleNumber)

	aroma := []string{}
	for i := 0; i < 3; i++ {
		aroma = append(aroma, "TEST_AROMA_"+strconv.Itoa(rand.Int()%aromaNumber))
	}

	imageURL := []string{}
	for i := 0; i < 5; i++ {
		imageURL = append(imageURL, "http:naver.com")
	}

	return beer.Beer{
		Name:      name,
		Brewery:   brewery,
		ABV:       abv,
		Country:   country,
		BeerStyle: beerStyle,
		Aroma:     aroma,
		ImageURL:  imageURL,
	}
}

func getRandomComment() beer.Comment {
	return beer.Comment{
		BeerID:  rand.Int63n(beerNumber) + 1,
		Content: "TEST_COMMENT_" + strconv.Itoa(rand.Int()),
		UserID:  rand.Int63n(beerNumber) + 1,
	}
}

func getRandomRate() beer.Rate {
	return beer.Rate{
		BeerID: rand.Int63n(beerNumber) + 1,
		Ratio:  rand.Float64()*rateBase + (5 - rateBase),
		UserID: rand.Int63n(beerNumber) + 1,
	}
}
