package beer_test

import (
	"fmt"
	"testing"

	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ucTestSuite))
}

type ucTestSuite struct {
	suite.Suite
	r  *mockBeerRepo
	uc beer.UseCase
}

func (ts *ucTestSuite) SetupTest() {
	ts.r = new(mockBeerRepo)
	ts.uc = beer.NewUseCase(ts.r)
}

func (ts *ucTestSuite) AfterTest() {
	ts.r.AssertExpectations(ts.T())
}

func (ts *ucTestSuite) Test_AddBeer() {
	var input beer.Beer
	faker.FakeData(&input)

	ts.r.On("AddBeer", input).Return(nil).Once()

	err := ts.uc.AddBeer(input)
	ts.NoError(err)
}

func (ts *ucTestSuite) Test_GetBeers() {
	var input beer.BeerQueryArgs
	faker.FakeData(&input)
	var output []beer.Beer
	faker.FakeData(&output)

	ts.r.On("GetBeers", input).Return(output, nil).Once()

	res, err := ts.uc.GetBeers(input)
	ts.NoError(err)
	ts.Equal(res, output)
}

func (ts *ucTestSuite) Test_GetBeer() {
	var input int64
	faker.FakeData(&input)
	var output *beer.Beer
	faker.FakeData(&output)

	ts.r.On("GetBeer", input).Return(output, nil).Once()

	res, err := ts.uc.GetBeer(input)
	ts.NoError(err)
	ts.Equal(res, output)
}

func (ts *ucTestSuite) Test_AddReview() {
	var nilBeer *beer.Beer

	ts.Run("Beer Finding Error", func() {
		var input beer.Review
		faker.FakeData(&input)

		ts.r.On("GetBeer", input.BeerID).Return(nilBeer, fmt.Errorf("finding beer error")).Once()

		err := ts.uc.AddReview(input)
		ts.Error(err)
	})

	ts.Run("No Beer Matching", func() {
		var input beer.Review
		faker.FakeData(&input)

		ts.r.On("GetBeer", input.BeerID).Return(nilBeer, nil).Once()

		err := ts.uc.AddReview(input)
		ts.Error(err)
	})
	ts.Run("First Review", func() {
		var input beer.Review
		faker.FakeData(&input)
		input.Ratio = 5.6
		var br beer.Beer
		faker.FakeData(&br)
		br.ID = input.BeerID
		br.ReviewCount = 0

		ts.r.On("GetBeer", input.BeerID).Return(&br, nil).Once()
		ts.r.On("UpdateBeerRateAvg", input.BeerID, input.Ratio).Return(nil).Once()
		ts.r.On("AddReview", input).Return(nil).Once()

		err := ts.uc.AddReview(input)
		ts.NoError(err)
	})
	ts.Run("None First Review - Pre Review Nil", func() {
		var input beer.Review
		faker.FakeData(&input)
		input.Ratio = 5.0
		var br beer.Beer
		faker.FakeData(&br)
		br.ID = input.BeerID
		br.ReviewCount = 4 // Not Zero
		br.RateAvg = 4.0

		var nilPreReview *beer.Review

		newRateAvg := 4.2 // (4.0 * 4 + 5.0) / 5

		ts.r.On("GetBeer", input.BeerID).Return(&br, nil).Once()
		ts.r.On("GetReviewByBeerIDAndUserID", input.BeerID, input.UserID).Return(nilPreReview, nil).Once()
		ts.r.On("UpdateBeerRateAvg", input.BeerID, newRateAvg).Return(nil).Once()
		ts.r.On("AddReview", input).Return(nil).Once()

		err := ts.uc.AddReview(input)
		ts.NoError(err)
	})
	ts.Run("None First Review - Pre Review Non Nil", func() {
		var input beer.Review
		faker.FakeData(&input)
		input.Ratio = 4.0
		var br beer.Beer
		faker.FakeData(&br)
		br.ID = input.BeerID
		br.ReviewCount = 5 // Not Zero
		br.RateAvg = 4.2

		var preReview *beer.Review
		faker.FakeData(&preReview)
		preReview.Ratio = 5

		newRateAvg := 4.0

		ts.r.On("GetBeer", input.BeerID).Return(&br, nil).Once()
		ts.r.On("GetReviewByBeerIDAndUserID", input.BeerID, input.UserID).Return(preReview, nil).Once()
		ts.r.On("UpdateBeerRateAvg", input.BeerID, newRateAvg).Return(nil).Once()
		ts.r.On("AddReview", input).Return(nil).Once()

		err := ts.uc.AddReview(input)
		ts.NoError(err)
	})
}

func (ts *ucTestSuite) Test_GetReviews() {
	var input int64
	faker.FakeData(&input)
	var output []beer.Review
	faker.FakeData(&output)

	ts.r.On("GetReviews", input).Return(output, nil).Once()

	res, err := ts.uc.GetReviews(input)
	ts.NoError(err)
	ts.Equal(res, output)
}

func (ts *ucTestSuite) Test_GetReviewsByUserID() {
	var input int64
	faker.FakeData(&input)
	var output []beer.Review
	faker.FakeData(&output)

	ts.r.On("GetReviewsByUserID", input).Return(output, nil).Once()

	res, err := ts.uc.GetReviewsByUserID(input)
	ts.NoError(err)
	ts.Equal(res, output)
}

func (ts *ucTestSuite) Test_GetReviewByBeerIDAndUserID() {
	var beerID, userID int64
	faker.FakeData(&beerID)
	faker.FakeData(&userID)
	var output *beer.Review
	faker.FakeData(&output)

	ts.r.On("GetReviewByBeerIDAndUserID", beerID, userID).Return(output, nil).Once()

	res, err := ts.uc.GetReviewByBeerIDAndUserID(beerID, userID)
	ts.NoError(err)
	ts.Equal(res, output)
}

func (ts *ucTestSuite) Test_GetRelatedBeers() {
	var beerID int64
	faker.FakeData(&beerID)
	var baseBeer *beer.Beer
	faker.FakeData(&baseBeer)

	var aromaRelatedBeers, styleRelatedBeers, randomlyRealtedBeers []beer.Beer
	faker.FakeData(&aromaRelatedBeers)
	faker.FakeData(&styleRelatedBeers)
	faker.FakeData(&randomlyRealtedBeers)

	var aromaRelatedQueryArgs, styleRelatedQueryArgs, randomlyRelatedQueryArgs beer.BeerQueryArgs
	aromaRelatedQueryArgs.Aroma = baseBeer.Aroma
	aromaRelatedQueryArgs.MaxCount = 5
	styleRelatedQueryArgs.BeerStyle = append(styleRelatedQueryArgs.BeerStyle, baseBeer.BeerStyle)
	styleRelatedQueryArgs.MaxCount = 5
	randomlyRelatedQueryArgs.MaxCount = 5

	var nilBeer *beer.Beer
	var nilBeers []beer.Beer

	ts.Run("GetBeer Error", func() {
		ts.r.On("GetBeer", beerID).Return(nilBeer, fmt.Errorf("beer not found")).Once()

		_, err := ts.uc.GetRelatedBeers(beerID)
		ts.Error(err)
	})
	ts.Run("AromaRelatedBeer Finding Error", func() {
		ts.r.On("GetBeer", beerID).Return(baseBeer, nil).Once()
		ts.r.On("GetBeers", aromaRelatedQueryArgs).Return(nilBeers, fmt.Errorf("beers not found")).Once()

		_, err := ts.uc.GetRelatedBeers(beerID)
		ts.Error(err)
	})
	ts.Run("StyleRelatedBeer Finding Error", func() {
		ts.r.On("GetBeer", beerID).Return(baseBeer, nil).Once()
		ts.r.On("GetBeers", aromaRelatedQueryArgs).Return(aromaRelatedBeers, nil).Once()
		ts.r.On("GetBeers", styleRelatedQueryArgs).Return(nilBeers, fmt.Errorf("beers not found")).Once()

		_, err := ts.uc.GetRelatedBeers(beerID)
		ts.Error(err)
	})
	ts.Run("RandomlyRelatedQueryArgs Finding Error", func() {
		ts.r.On("GetBeer", beerID).Return(baseBeer, nil).Once()
		ts.r.On("GetBeers", aromaRelatedQueryArgs).Return(aromaRelatedBeers, nil).Once()
		ts.r.On("GetBeers", styleRelatedQueryArgs).Return(styleRelatedBeers, nil).Once()
		ts.r.On("GetBeers", randomlyRelatedQueryArgs).Return(nilBeers, fmt.Errorf("beers not found")).Once()

		_, err := ts.uc.GetRelatedBeers(beerID)
		ts.Error(err)
	})
	ts.Run("Success", func() {
		ts.r.On("GetBeer", beerID).Return(baseBeer, nil).Once()
		ts.r.On("GetBeers", aromaRelatedQueryArgs).Return(aromaRelatedBeers, nil).Once()
		ts.r.On("GetBeers", styleRelatedQueryArgs).Return(styleRelatedBeers, nil).Once()
		ts.r.On("GetBeers", randomlyRelatedQueryArgs).Return(randomlyRealtedBeers, nil).Once()

		_, err := ts.uc.GetRelatedBeers(beerID)
		ts.NoError(err)
		// Output check is omitted. For it's randomness in beer order
	})
}

type mockBeerRepo struct {
	mock.Mock
}

func (r *mockBeerRepo) AddBeer(beer beer.Beer) error {
	ret := r.Called(beer)
	return ret.Error(0)
}

func (r *mockBeerRepo) GetBeers(args beer.BeerQueryArgs) ([]beer.Beer, error) {
	ret := r.Called(args)
	return ret.Get(0).([]beer.Beer), ret.Error(1)
}

func (r *mockBeerRepo) GetBeer(beerID int64) (*beer.Beer, error) {
	ret := r.Called(beerID)
	return ret.Get(0).(*beer.Beer), ret.Error(1)
}

func (r *mockBeerRepo) UpdateBeerRateAvg(beerID int64, rateAvg float64) error {
	ret := r.Called(beerID, rateAvg)
	return ret.Error(0)
}

func (r *mockBeerRepo) AddReview(review beer.Review) error {
	ret := r.Called(review)
	return ret.Error(0)
}

func (r *mockBeerRepo) GetReviews(beerID int64) ([]beer.Review, error) {
	ret := r.Called(beerID)
	return ret.Get(0).([]beer.Review), ret.Error(1)
}

func (r *mockBeerRepo) GetReviewsByUserID(userID int64) ([]beer.Review, error) {
	ret := r.Called(userID)
	return ret.Get(0).([]beer.Review), ret.Error(1)
}

func (r *mockBeerRepo) GetReviewCount(beerID int64) (int64, error) {
	ret := r.Called(beerID)
	return ret.Get(0).(int64), ret.Error(1)
}

func (r *mockBeerRepo) GetReviewByBeerIDAndUserID(beerID int64, userID int64) (*beer.Review, error) {
	ret := r.Called(beerID, userID)
	return ret.Get(0).(*beer.Review), ret.Error(1)
}

func (r *mockBeerRepo) AddFavorite(favorite beer.Favorite) error {
	ret := r.Called(favorite)
	return ret.Error(0)
}

func (r *mockBeerRepo) GetFavorites(userID int64) ([]beer.Favorite, error) {
	ret := r.Called(userID)
	return ret.Get(0).([]beer.Favorite), ret.Error(1)
}

func (r *mockBeerRepo) AddUserBeerConfig(userBeerConfig beer.UserBeerConfig) error {
	ret := r.Called(userBeerConfig)
	return ret.Error(0)
}

func (r *mockBeerRepo) GetUserBeerConfig(userID int64) (*beer.UserBeerConfig, error) {
	ret := r.Called(userID)
	return ret.Get(0).(*beer.UserBeerConfig), ret.Error(1)
}
