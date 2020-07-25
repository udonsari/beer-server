package datasource

type DBBeer struct {
	ID        int `orm:"auto"`
	Name      string
	Brewery   string
	ABV       float64
	Country   string
	BeerStyle string
	Aroma     string
}
