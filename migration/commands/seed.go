package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/UdonSari/beer-server/domain/beer"
	beerRepo "github.com/UdonSari/beer-server/domain/beer/repo"
	"github.com/UdonSari/beer-server/main/server"
	"github.com/davecgh/go-spew/spew"
	"github.com/urfave/cli/v2"
)

// 여전히 향과 사진들은 찾기 어렵다
var fakeAromaList []string

type seedCommand struct {
	d          *server.Dependency
	hostStatic string
	mapper
}

func NewSeedCommand(d *server.Dependency) *seedCommand {
	rand.Seed(time.Now().UnixNano())
	hostStatic := fmt.Sprintf("%s:%s/static", d.Host(), d.PortStr())

	// From https://winning-homebrew.com/beer-flavor-descriptors.html
	// fakeAromaList = []string{
	// 	"malty", "caramel", "roast", "coffee", "grass", "banana", "apple", "peach", "mango", "orange", "grapefruit", "vinegar", "nutty",
	// }

	return &seedCommand{
		d:          d,
		hostStatic: hostStatic,
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
	br := beerRepo.New(c.d.MysqlDB(false), 100)
	bu := beer.NewUseCase(br)

	var dataBeers []Beer
	// rawBeers, err := ioutil.ReadFile("../beerdata/open-beer-database.json")
	rawBeers, err := ioutil.ReadFile("/src/beer-server/migration/beerdata/open-beer-database.json")
	if err != nil {
		log.Fatalf("read err %+v", err)
	}
	err = json.Unmarshal(rawBeers, &dataBeers)
	if err != nil {
		log.Fatalf("unmarshal err %+v", err)
	}

	beerCount := 0
	beerStyleList := make(map[string]int)
	for i, dataBeer := range dataBeers {
		if !c.isValidDataBeer(dataBeer) {
			continue
		}
		beer := c.mapper.MapDataBeerToBeer(dataBeer)

		// Override
		beer.ImageURL = []string{fmt.Sprintf("%s/%s", c.hostStatic, "basic_beer_image.png")}
		beer.ThumbnailImage = fmt.Sprintf("%s/%s", c.hostStatic, "basic_beer_image.png")

		log.Printf("trying to put %vth beer %v", i, beer)
		if err := bu.AddBeer(beer); err != nil {
			log.Fatalf("failed to add %+v with err %+v", spew.Sdump(beer), spew.Sdump(err))
		}
		beerCount++

		// Temporal
		beerStyleList[dataBeer.Fields.BeerStyle] = 1
	}
	log.Printf("%v Beer is successfuly seeded !", beerCount)
	log.Printf("%v Styles !", len(beerStyleList))
	for beerStyle := range beerStyleList {
		fmt.Printf("%+v, ", beerStyle)
	}
	/*
		*** Beer Count ***
		2414

		*** Style List ***

		Specialty Beer, Porter, American-Style Stout, Winter Warmer, Foreign (Export)-Style Stout, Fruit Beer, Belgian-Style Pale Strong Ale, English-Style Pale Mild Ale, German-Style Pilsener, Out of Category, American-Style Light Lager, German-Style Heller Bock/Maibock, Scotch Ale, German-Style Brown Ale/Altbier, French & Belgian-Style Saison, Belgian-Style Dubbel, Pumpkin Beer, Belgian-Style Dark Strong Ale, Classic English-Style Pale Ale, Imperial or Double Red Ale, Dark American-Belgo-Style Ale, American Rye Ale or Lager, Belgian-Style Pale Ale, American-Style Brown Ale, English-Style India Pale Ale, Golden or Blonde Ale, Classic Irish-Style Dry Stout, Special Bitter or Best Bitter, American-Style Pale Ale, Light American Wheat Ale or Lager, American-Style Imperial Stout, Strong Ale, Kellerbier - Ale, Traditional German-Style Bock, Baltic-Style Porter, American-Style Cream Ale or Lager, English-Style Dark Mild Ale, Sweet Stout, South German-Style Hefeweizen, American-Style India Black Ale, Ordinary Bitter, American-Style India Pale Ale, Herb and Spice Beer, American-Style Barley Wine Ale, American-Style Strong Pale Ale, Scottish-Style Light Ale, German-Style Schwarzbier, German-Style Doppelbock, Irish-Style Red Ale, Belgian-Style Quadrupel, Extra Special Bitter, Belgian-Style Tripel, American-Style Dark Lager, American-Style Amber/Red Ale, Oatmeal Stout, Imperial or Double India Pale Ale, South German-Style Weizenbock, Belgian-Style White, Old Ale, Belgian-Style Fruit Lambic, American-Style Lager, Other Belgian-Style Ales, German-Style Oktoberfest
	*/
	return nil
}

func (c *seedCommand) isValidDataBeer(dataBeer Beer) bool {
	return dataBeer.Fields.ABV > 1.0 &&
		dataBeer.Fields.Brewery != "" &&
		dataBeer.Fields.Country != "" &&
		dataBeer.Fields.BeerStyle != ""
}

type mapper struct{}

func (m *mapper) MapDataBeerToBeer(dataBeer Beer) beer.Beer {
	aroma := []string{}
	imageURL := []string{}
	thumbnailImage := ""

	return beer.Beer{
		Name:      dataBeer.Fields.Name,
		Brewery:   dataBeer.Fields.Brewery,
		ABV:       dataBeer.Fields.ABV,
		Country:   dataBeer.Fields.Country,
		BeerStyle: dataBeer.Fields.BeerStyle,

		// Still Fake
		Aroma:          aroma,
		ImageURL:       imageURL,
		ThumbnailImage: thumbnailImage,
	}
}

type Beer struct {
	Fields Field `json:"fields"`
}

type Field struct {
	// Only care data what I want
	Name      string  `json:"name"`
	Brewery   string  `json:"name_breweries"`
	ABV       float64 `json:"abv"`
	Country   string  `json:"country"`
	BeerStyle string  `json:"style_name"`
}

/*
Example

      "datasetid":"open-beer-database@public-us",
      "recordid":"21427b5076ea6e6adf0f997b460f0c822b0dcdc9",
      "fields":{
         "brewery_id":"842",
         "city":"Mill Creek",
         "name":"Porter",
         "cat_name":"Irish Ale",
         "country":"United States",
         "cat_id":"2",
         "upc":0,
         "coordinates":[
            47.8774,
            -122.211
         ],
         "srm":0,
         "last_mod":"2010-07-23T05:00:00+09:00",
         "state":"Washington",
         "add_user":"0",
         "abv":0.0,
         "address1":"13300 Bothell-Everett Highway #304",
         "name_breweries":"McMenamins Mill Creek",
         "style_name":"Porter",
         "id":"716",
         "ibu":0,
         "style_id":"25"
      },
      "geometry":{
         "type":"Point",
         "coordinates":[
            -122.211,
            47.8774
         ]
      },
      "record_timestamp":"2016-09-26T13:21:38.074+09:00"
   },
*/
