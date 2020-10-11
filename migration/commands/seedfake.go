package commands

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/UdonSari/beer-server/domain/beer"
	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	"github.com/UdonSari/beer-server/domain/user"
	userRepo "github.com/UdonSari/beer-server/domain/user/repo"
	"github.com/UdonSari/beer-server/main/server"
	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
)

const (
	reviewNumber    = 1000
	beerNumber      = 100
	userNumber      = 100
	rateBase        = 3
	abvLimit        = 10
	countryNumber   = 10
	beerStyleNumber = 5
	breweryNumber   = 100
	aromaNumber     = 5

	imageWidth  = 320
	imageHeight = 480
)

type seedFakeCommand struct {
	d *server.Dependency
}

func NewSeedFakeCommand(d *server.Dependency) *seedFakeCommand {
	rand.Seed(time.Now().UnixNano())
	return &seedFakeCommand{
		d: d,
	}
}

func (c *seedFakeCommand) Command() *cli.Command {
	return &cli.Command{
		Name:   "seed-fake",
		Usage:  "Used to seed fake data used by server",
		Action: c.main,
	}
}

func (c *seedFakeCommand) main(ctx *cli.Context) error {
	br := beerRepo.New(c.d.MysqlDB(false), 100)
	ur := userRepo.New(c.d.MysqlDB(false))
	bu := beer.NewUseCase(br)
	uu := user.NewUseCase(ur, "", "", "")

	for i := 0; i < beerNumber; i++ {
		beer := getRandomBeer()
		log.Printf("trying to put %vth beer %v", i, beer)
		if err := bu.AddBeer(beer); err != nil {
			log.Fatalf("failed to add %+v with err %+v", spew.Sdump(beer), spew.Sdump(err))
		}
	}

	for i := 0; i < reviewNumber; i++ {
		review := getRandomReview()
		log.Printf("trying to put %vth review %v", i, review)
		if err := bu.AddReview(review); err != nil {
			// Review가 duplicate으로 들어ㅏ가지 않을 수 있다. 해당 경우 무시.
			log.Printf("failed to add %+v with err %+v", spew.Sdump(review), spew.Sdump(err))
			i--
			continue
		}
	}

	for i := 0; i < userNumber; i++ {
		user := getRandomUser()
		log.Printf("trying to put %vth user %v", i, user)
		if err := uu.CreateUser(user); err != nil {
			// User duplicate으로 들어ㅏ가지 않을 수 있다. 해당 경우 무시.
			log.Printf("failed to add %+v with err %+v", spew.Sdump(user), spew.Sdump(err))
			i--
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
		newAroma := ""
		for true {
			notDuplicatedCount := 0
			newAroma = "TEST_AROMA_" + strconv.Itoa(rand.Int()%aromaNumber)
			for ; notDuplicatedCount < i; notDuplicatedCount++ {
				if newAroma == aroma[notDuplicatedCount] {
					break
				}
			}

			if notDuplicatedCount == i {
				break
			}
		}
		aroma = append(aroma, newAroma)
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
		UserID:  rand.Int63n(userNumber) + 1,
	}
}

func getRandomUser() user.User {
	return user.User{
		ExternalID: "TEST_EXTERNAL_ID_" + strconv.Itoa(rand.Int()),
		Properties: user.Properties{
			NickName:       "TEST_NICKNAME_" + strconv.Itoa(rand.Int()),
			ProfileImage:   fmt.Sprintf("https://picsum.photos/%v/%v", imageWidth, imageHeight),
			ThumbnailImage: fmt.Sprintf("https://picsum.photos/%v/%v", imageWidth, imageHeight),
		},
	}
}
