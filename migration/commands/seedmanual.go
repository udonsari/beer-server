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
		// {
		// 	Name:           "봄마실",
		// 	Brewery:        "맥파이",
		// 	ABV:            4.0,
		// 	Country:        "Korea",
		// 	BeerStyle:      "Saison",
		// 	Aroma:          []string{"Malty", "Spicy"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_bringspring.jpeg")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_bringspring.jpeg"),
		// },
		// {
		// 	Name:           "슬리퍼",
		// 	Brewery:        "맥파이",
		// 	ABV:            11.0,
		// 	Country:        "Korea",
		// 	BeerStyle:      "Quadrupel",
		// 	Aroma:          []string{"Malty", "Spicy", "Banana"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_sleeper.jpg")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_sleeper.jpg"),
		// },
		// {
		// 	Name:           "겨울동지",
		// 	Brewery:        "맥파이",
		// 	ABV:            3.8,
		// 	Country:        "Korea",
		// 	BeerStyle:      "Brown Ale",
		// 	Aroma:          []string{"Nutty", "Caramel"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_wintersouls.jpg")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_wintersouls.jpg"),
		// },
		// {
		// 	Name:           "가을가득",
		// 	Brewery:        "맥파이",
		// 	ABV:            5.5,
		// 	Country:        "Korea",
		// 	BeerStyle:      "Rye Amber",
		// 	Aroma:          []string{"Spicy", "Orange"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_fulloffall.jpg")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_fulloffall.jpg"),
		// },
		// {
		// 	Name:           "동네친구",
		// 	Brewery:        "맥파이",
		// 	ABV:            5.0,
		// 	Country:        "Korea",
		// 	BeerStyle:      "Pilsener",
		// 	Aroma:          []string{"Malty", "Grass"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_oldpals.png")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_oldpals.png"),
		// },
		// {
		// 	Name:           "고스트",
		// 	Brewery:        "맥파이",
		// 	ABV:            4.5,
		// 	Country:        "Korea",
		// 	BeerStyle:      "Gose",
		// 	Aroma:          []string{"Orange", "Vinegar"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_ghost.jpg")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_ghost.jpg"),
		// },
		// {
		// 	Name:           "맥파이 IPA",
		// 	Brewery:        "맥파이",
		// 	ABV:            6.5,
		// 	Country:        "Korea",
		// 	BeerStyle:      "c",
		// 	Aroma:          []string{"Orange"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "magpie_indiapaleale.png")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "magpie_indiapaleale.png"),
		// },
		// {
		// 	Name:           "맥파이 KÖLSCH",
		// 	Brewery:        "맥파이",
		// 	ABV:            4.8,
		// 	Country:        "Korea",
		// 	BeerStyle:      "Kolsch",
		// 	Aroma:          []string{"Malty"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "magpie_kolsch.png")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "magpie_kolsch.png"),
		// },
		// {
		// 	Name:           "맥파이 Porter",
		// 	Brewery:        "맥파이",
		// 	ABV:            5.6,
		// 	Country:        "Korea",
		// 	BeerStyle:      "Porter",
		// 	Aroma:          []string{"Malty", "Roast", "Coffee"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "magpie_porter.png")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "magpie_porter.png"),
		// },
		// {
		// 	Name:           "맥파이 Pale Ale",
		// 	Brewery:        "맥파이",
		// 	ABV:            4.8,
		// 	Country:        "Korea",
		// 	BeerStyle:      "Pale Ale",
		// 	Aroma:          []string{"Orange", "Mango"},
		// 	ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "macpie_paleale.png")},
		// 	ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "macpie_paleale.png"),
		// },
		{
			Name:           "ㅋ IPA",
			Brewery:        "더부스",
			ABV:            4.5,
			Country:        "Korea",
			BeerStyle:      "India Pale Ale",
			Aroma:          []string{"Peach", "Mango"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_kieukipa.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_kieukipa.png"),
		},
		{
			Name:           "대강 페일 에일",
			Brewery:        "더부스",
			ABV:            4.6,
			Country:        "Korea",
			BeerStyle:      "Pale Ale",
			Aroma:          []string{"Orange"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_kieukipa.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_taegangpaleale.png"),
		},
		{
			Name:           "대강 페일 에일",
			Brewery:        "더부스",
			ABV:            4.6,
			Country:        "Korea",
			BeerStyle:      "Pale Ale",
			Aroma:          []string{"Orange"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_kieukipa.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_taegangpaleale.png"),
		},
		{
			Name:           "윗 마이 엑스",
			Brewery:        "더부스",
			ABV:            5.5,
			Country:        "Korea",
			BeerStyle:      "Witbier",
			Aroma:          []string{"Orange", "Spicy", "Grass"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_witmyex.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_witmyex.png"),
		},
		{
			Name:           "국민 IPA",
			Brewery:        "더부스",
			ABV:            6.5,
			Country:        "Korea",
			BeerStyle:      "India Pale Ale",
			Aroma:          []string{"Orange", "Pineapple", "Melon"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_kookminipa.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_kookminipa.png"),
		},
		{
			Name:           "경리단 힙스터",
			Brewery:        "더부스",
			ABV:            4.5,
			Country:        "Korea",
			BeerStyle:      "India Pale Ale",
			Aroma:          []string{"Orange", "Grass"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_gyunglidanhipster.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_gyunglidanhipster.png"),
		},
		{
			Name:           "긍정신 레드 에일",
			Brewery:        "더부스",
			ABV:            5.0,
			Country:        "Korea",
			BeerStyle:      "Red Ale",
			Aroma:          []string{"Malty", "Caramel"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_thegreatgodoffun.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_thegreatgodoffun.png"),
		},
		{
			Name:           "치믈리에일",
			Brewery:        "더부스",
			ABV:            5.0,
			Country:        "Korea",
			BeerStyle:      "Pale Ale",
			Aroma:          []string{"Pineapple", "Orange"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_chimmeliale.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_chimmeliale.png"),
		},
		{
			Name:           "LGBTQ Smoothie IPA",
			Brewery:        "더부스",
			ABV:            7.5,
			Country:        "Korea",
			BeerStyle:      "India Pale Ale",
			Aroma:          []string{"Blackberry"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_lgbtqsmoothieipa.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_lgbtqsmoothieipa.png"),
		},
		{
			Name:           "유레카 서울",
			Brewery:        "더부스",
			ABV:            6.5,
			Country:        "Korea",
			BeerStyle:      "ETC",
			Aroma:          []string{"Peach"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_eurekaseoul.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "thebooth_eurekaseoul.png"),
		},
		{
			Name:           "LIFE IPA 마릴린먼로",
			Brewery:        "크래프트브로스",
			ABV:            6.5,
			Country:        "Korea",
			BeerStyle:      "New England IPA",
			Aroma:          []string{"Orange", "Pineapple"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "craftbros_lifeipamarilynmonroe.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "craftbros_lifeipamarilynmonroe.png"),
		},
		{
			Name:           "LIFE IPA 체게바라",
			Brewery:        "크래프트브로스",
			ABV:            6.5,
			Country:        "Korea",
			BeerStyle:      "New England IPA",
			Aroma:          []string{"Orange", "Pineapple"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "craftbros_lifeipacheguevara.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "craftbros_lifeipacheguevara.png"),
		},
		{
			Name:           "홉스플래쉬 IPA",
			Brewery:        "플레이그라운드",
			ABV:            6.7,
			Country:        "Korea",
			BeerStyle:      "New England IPA",
			Aroma:          []string{"Orange", "Pineapple"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_hopsplash.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_hopsplash.png"),
		},
		{
			Name:           "홉스플래쉬 IPA",
			Brewery:        "플레이그라운드",
			ABV:            6.7,
			Country:        "Korea",
			BeerStyle:      "New England IPA",
			Aroma:          []string{"Orange", "Pineapple"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_hopsplash.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_hopsplash.png"),
		},
		{
			Name:           "스칼라 벨지안 블론드 에일",
			Brewery:        "플레이그라운드",
			ABV:            5.4,
			Country:        "Korea",
			BeerStyle:      "ETC",
			Aroma:          []string{"Malty", "Caramel"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_scala.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_scala.png"),
		},
		{
			Name:           "위치 초콜릿 스타우트",
			Brewery:        "플레이그라운드",
			ABV:            5.7,
			Country:        "Korea",
			BeerStyle:      "Stout",
			Aroma:          []string{"Malty", "Caramel", "Roast", "Coffee", "Chocolate"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_witch.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_witch.png"),
		},
		{
			Name:           "헌치백 세션 IPA",
			Brewery:        "플레이그라운드",
			ABV:            4.0,
			Country:        "Korea",
			BeerStyle:      "India Pale Ale",
			Aroma:          []string{"Orange"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_hunchback.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_hunchback.png"),
		},
		{
			Name:           "조커 골든 페일 에일",
			Brewery:        "플레이그라운드",
			ABV:            5.6,
			Country:        "Korea",
			BeerStyle:      "Pale Ale",
			Aroma:          []string{"Orange"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_joker.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_joker.png"),
		},
		{
			Name:           "마담 체리 위트 에일",
			Brewery:        "플레이그라운드",
			ABV:            5.6,
			Country:        "Korea",
			BeerStyle:      "Witbier",
			Aroma:          []string{"Cherry"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_madame.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_madame.png"),
		},
		{
			Name:           "미스트리스 사워 에일",
			Brewery:        "플레이그라운드",
			ABV:            5.4,
			Country:        "Korea",
			BeerStyle:      "Sour Ale",
			Aroma:          []string{"Lemon", "Pineapple"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_mistress.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_mistress.png"),
		},
		{
			Name:           "몽크 IPA",
			Brewery:        "플레이그라운드",
			ABV:            7.2,
			Country:        "Korea",
			BeerStyle:      "India Pale Ale",
			Aroma:          []string{"Malty", "Pineapple"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_monk.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_monk.png"),
		},
		{
			Name:           "젠틀맨 라거",
			Brewery:        "플레이그라운드",
			ABV:            7.6,
			Country:        "Korea",
			BeerStyle:      "Pilsener",
			Aroma:          []string{"Malty"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "playground_gentleman.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "playground_gentleman.png"),
		},
		{
			Name:           "Fandango",
			Brewery:        "Toppling Goliath",
			ABV:            4.0,
			Country:        "USA",
			BeerStyle:      "Sour Ale",
			Aroma:          []string{"Mango", "Passion Fruit"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_dragonfandango.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_dragonfandango.png"),
		},
		{
			Name:           "Pseudo Sue",
			Brewery:        "Toppling Goliath",
			ABV:            5.8,
			Country:        "USA",
			BeerStyle:      "Pale Ale",
			Aroma:          []string{"Grapefruit", "Mango", "Grass"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_pseudosue.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_pseudosue.png"),
		},
		{
			Name:           "Mornin Latte",
			Brewery:        "Toppling Goliath",
			ABV:            8.9,
			Country:        "USA",
			BeerStyle:      "Stout",
			Aroma:          []string{"Malty", "Roast", "Coffee"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_morninlatte.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_morninlatte.png"),
		},
		{
			Name:           "King Sue",
			Brewery:        "Toppling Goliath",
			ABV:            7.8,
			Country:        "USA",
			BeerStyle:      "India Pale Ale",
			Aroma:          []string{"Orange", "Mango", "Pineapple", "Grapefruit"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_kingsue.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_kingsue.png"),
		},
		{
			Name:           "Dorothy's",
			Brewery:        "Toppling Goliath",
			ABV:            5.5,
			Country:        "USA",
			BeerStyle:      "ETC",
			Aroma:          []string{"Malty"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_dorothys.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_dorothys.png"),
		},
		{
			Name:           "Pompeii",
			Brewery:        "Toppling Goliath",
			ABV:            5.8,
			Country:        "USA",
			BeerStyle:      "India Pale Ale",
			Aroma:          []string{"Pineapple", "Mango"},
			ImageURL:       []string{fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_pompeii.png")},
			ThumbnailImage: fmt.Sprintf("%s/%s", c.hostStatic, "topplinggoliath_pompeii.png"),
		},
	}
}
