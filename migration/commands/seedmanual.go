package commands

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/UdonSari/beer-server/domain/beer"
	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	"github.com/UdonSari/beer-server/main/server"
	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
)

type seedManualCommand struct {
	d          *server.Dependency
	hostStatic string
}

func NewSeedManualCommand(d *server.Dependency) *seedManualCommand {
	rand.Seed(time.Now().UnixNano())
	hostStatic := fmt.Sprintf("%s:%s/static", d.Host(), d.PortStr())
	return &seedManualCommand{
		d:          d,
		hostStatic: hostStatic,
	}
}

func (c *seedManualCommand) Command() *cli.Command {
	return &cli.Command{
		Name:   "seed-manual",
		Usage:  "Used to seed manual data used by server",
		Action: c.main,
	}
}

func (c *seedManualCommand) main(ctx *cli.Context) error {
	br := beerRepo.New(c.d.MysqlDB(false), 100)
	bu := beer.NewUseCase(br)

	beers := c.GetBeers()
	for _, beer := range beers {
		if err := bu.AddBeer(beer); err != nil {
			log.Fatalf("failed to add %+v with err %+v", spew.Sdump(beer), spew.Sdump(err))
		}
	}
	log.Printf("%v Beer is successfuly seeded !", len(beers))
	return nil
}

func (c *seedManualCommand) GetBeers() []beer.Beer {
	return []beer.Beer{
		{
			Name:           "봄마실",
			Brewery:        "맥파이",
			ABV:            4.0,
			Country:        "Korea",
			BeerStyle:      "Saison",
			Aroma:          []string{"Malty", "Spicy"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_bringspring.jpeg")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_bringspring.jpeg"),
		},
		{
			Name:           "슬리퍼",
			Brewery:        "맥파이",
			ABV:            11.0,
			Country:        "Korea",
			BeerStyle:      "Quadrupel",
			Aroma:          []string{"Malty", "Spicy", "Banana"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_sleeper.jpg")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_sleeper.jpg"),
		},
		{
			Name:           "겨울동지",
			Brewery:        "맥파이",
			ABV:            3.8,
			Country:        "Korea",
			BeerStyle:      "Brown Ale",
			Aroma:          []string{"Nutty", "Caramel"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_wintersouls.jpg")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_wintersouls.jpg"),
		},
		{
			Name:           "가을가득",
			Brewery:        "맥파이",
			ABV:            5.5,
			Country:        "Korea",
			BeerStyle:      "Rye Amber",
			Aroma:          []string{"Spicy", "Orange"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_fulloffall.jpg")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_fulloffall.jpg"),
		},
		{
			Name:           "동네친구",
			Brewery:        "맥파이",
			ABV:            5.0,
			Country:        "Korea",
			BeerStyle:      "Pilsener",
			Aroma:          []string{"Malty", "Grass"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_oldpals.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_oldpals.png"),
		},
		{
			Name:           "고스트",
			Brewery:        "맥파이",
			ABV:            4.5,
			Country:        "Korea",
			BeerStyle:      "Gose",
			Aroma:          []string{"Orange", "Vinegar"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_ghost.jpg")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_ghost.jpg"),
		},
		{
			Name:           "맥파이 IPA",
			Brewery:        "맥파이",
			ABV:            6.5,
			Country:        "Korea",
			BeerStyle:      "India Pale Ale",
			Aroma:          []string{"Orange"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "magpie_indiapaleale.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "magpie_indiapaleale.png"),
		},
		{
			Name:           "맥파이 KÖLSCH",
			Brewery:        "맥파이",
			ABV:            4.8,
			Country:        "Korea",
			BeerStyle:      "Kolsch",
			Aroma:          []string{"Malty"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "magpie_kolsch.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "magpie_kolsch.png"),
		},
		{
			Name:           "맥파이 Porter",
			Brewery:        "맥파이",
			ABV:            5.6,
			Country:        "Korea",
			BeerStyle:      "Porter",
			Aroma:          []string{"Malty", "Roast", "Coffee"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "magpie_porter.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "magpie_porter.png"),
		},
		{
			Name:           "맥파이 Pale Ale",
			Brewery:        "맥파이",
			ABV:            4.8,
			Country:        "Korea",
			BeerStyle:      "Pale Ale",
			Aroma:          []string{"Orange", "Mango"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_paleale.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_paleale.png"),
		},
	}
}
