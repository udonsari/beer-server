package commands

import (
	"fmt"
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
	reviewNumber    = 1000
	beerNumber      = 100
	rateBase        = 3
	abvLimit        = 10
	countryNumber   = 10
	beerStyleNumber = 5
	breweryNumber   = 100
	aromaNumber     = 5

	imageWidth  = 320
	imageHeight = 480
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

	for i := 0; i < reviewNumber; i++ {
		review := getRandomReview()
		log.Printf("trying to put %vth review %v", i, review)
		if err := bu.AddReview(review); err != nil {
			// Review가 duplicate으로 들어ㅏ가지 않을 수 있다. 해당 경우 무시.
			log.Printf("failed to add %+v with err %+v", spew.Sdump(review), spew.Sdump(err))
			continue
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
		imageURL = append(imageURL, fmt.Sprintf("https://picsum.photos/%v/%v", imageWidth, imageHeight))
	}

	thumbnailImage := fmt.Sprintf("https://picsum.photos/%v/%v", imageWidth, imageHeight)

	return beer.Beer{
		Name:           name,
		Brewery:        brewery,
		ABV:            abv,
		Country:        country,
		BeerStyle:      beerStyle,
		Aroma:          aroma,
		ImageURL:       imageURL,
		ThumbnailImage: thumbnailImage,
	}
}

func getRandomReview() beer.Review {
	return beer.Review{
		BeerID:  rand.Int63n(beerNumber) + 1,
		Content: "TEST_CONTENT_" + strconv.Itoa(rand.Int()),
		Ratio:   rand.Float64()*rateBase + (5 - rateBase),
		UserID:  rand.Int63n(beerNumber) + 1,
	}
}
