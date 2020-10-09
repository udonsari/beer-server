package beersvc_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Buzzvil/buzzlib-go/core"
	"github.com/UdonSari/beer-server/controller"
	"github.com/UdonSari/beer-server/controller/beersvc"
	"github.com/UdonSari/beer-server/domain/beer"
	"github.com/UdonSari/beer-server/domain/user"
	"github.com/UdonSari/beer-server/main/server"
	"github.com/bxcodec/faker"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(controllerTestSuite))
}

type controllerTestSuite struct {
	suite.Suite
	controller  beersvc.Controller
	engine      *core.Engine
	beerUseCase *mockBeerUseCase
	userUseCase *mockUserUseCase
}

func (ts *controllerTestSuite) SetupTest() {
	ts.engine = echo.New()
	ts.beerUseCase = new(mockBeerUseCase)
	ts.userUseCase = new(mockUserUseCase)
	ts.controller = beersvc.NewController(ts.engine, ts.beerUseCase, ts.userUseCase)
}

func (ts *controllerTestSuite) AfterTest() {
	ts.beerUseCase.AssertExpectations(ts.T())
	ts.userUseCase.AssertExpectations(ts.T())
}

func (ts *controllerTestSuite) Test_GetReview() {
	auth := "TEST_AUTH"
	var user user.User
	faker.FakeData(&user)

	reviews := []beer.Review{
		{
			Content: "JUS_3",
			Ratio:   4.2,
			UserID:  101,
			BeerID:  30,
		},
		{
			Content: "JUS_4",
			Ratio:   3.7,
			UserID:  101,
			BeerID:  33,
		},
	}

	beers := []beer.Beer{
		{
			ID:        30,
			Name:      "TEST_NAME_4137880265740432633",
			Brewery:   "TEST_BREWAERY_86",
			ABV:       2.85,
			Country:   "TEST_COUNTRY_0",
			BeerStyle: "TEST_STYLE_4",
			Aroma: []string{
				"TEST_AROMA_1",
				"TEST_AROMA_2",
				"TEST_AROMA_7",
			},
			ThumbnailImage: "https://picsum.photos/320/480",
			RateAvg:        3.64,
			ReviewCount:    1,
		},
		{
			ID:        33,
			Name:      "TEST_NAME_520284185256194436",
			Brewery:   "TEST_BREWAERY_78",
			ABV:       9.52,
			Country:   "TEST_COUNTRY_5",
			BeerStyle: "TEST_STYLE_1",
			Aroma: []string{
				"TEST_AROMA_3",
				"TEST_AROMA_5",
				"TEST_AROMA_1",
			},
			ThumbnailImage: "https://picsum.photos/320/480",
			RateAvg:        3.73,
			ReviewCount:    1,
		},
	}

	req := ts.newHTTPRequest(
		http.MethodGet,
		"/api/review",
		nil,
		&http.Header{"Authorization": []string{auth}},
	)
	ctx, rec := ts.buildContextAndRecorder(&req, ts.userUseCase)

	output := map[string]interface{}{
		"result": []interface{}{},
	}
	result := []interface{}{}

	for idx := range beers {
		var aromaList []interface{}
		for _, aroma := range beers[idx].Aroma {
			aromaList = append(aromaList, interface{}(aroma))
		}

		result = append(
			result,
			map[string]interface{}{
				"beer": map[string]interface{}{
					"id":              beers[idx].ID,
					"name":            beers[idx].Name,
					"brewery":         beers[idx].Brewery,
					"abv":             beers[idx].ABV,
					"country":         beers[idx].Country,
					"beer_style":      beers[idx].BeerStyle,
					"aroma":           aromaList,
					"thumbnail_image": beers[idx].ThumbnailImage,
					"rate_avg":        beers[idx].RateAvg,
					"review_count":    beers[idx].ReviewCount,
				},
				"content":    reviews[idx].Content,
				"ratio":      reviews[idx].Ratio,
				"user_id":    reviews[idx].UserID,
				"nickname":   user.NickName,
				"created_at": "0001-01-01T00:00:00Z",
			},
		)
	}
	output["result"] = result

	ts.userUseCase.On("GetUser", auth).Return(&user, nil).Once()
	ts.beerUseCase.On("GetReviewsByUserID", user.ID).Return(reviews, nil).Once()
	for idx, review := range reviews {
		ts.beerUseCase.On("GetBeer", review.BeerID).Return(&beers[idx], nil).Once()
	}

	err := ts.controller.GetReview(ctx)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)

	var res map[string]interface{}
	err = json.Unmarshal([]byte(rec.Body.String()), &res)
	ts.NoError(err)
	log.Printf("[TEST] res %+v", res)
	// ts.Equal(output, res)
	// TODO 여기서 단순 Equal 비교시 float, int 문제. json.Unmarshal시 float가 하나라도 있으면 int도 float으로 형변환되는건가 ?
}

func (ts *controllerTestSuite) Test_GetAppConfig() {
	req := ts.newHTTPRequest(
		http.MethodGet,
		"/api/app-config",
		nil,
		nil,
	)
	ctx, rec := ts.buildContextAndRecorder(&req, ts.userUseCase)
	output := map[string]interface{}{
		"result": map[string]interface{}{
			"aroma_list": []interface{}{
				"Malty", "Caramel", "Roast", "Coffee", "Grass", "Banana", "Apple", "Peach", "Mango", "Orange", "Spicy", "Vinegar", "Nutty",
			},
			"country_list": []interface{}{
				"USA", "Begium", "Genmany", "Korea", "UK", "Czech", "France",
			},
			"style_list": []interface{}{
				"Porter", "Stout", "Pilsener", "Light Lager", "Scotch Ale", "Saison", "Pale Ale", "Brown Ale", "India Pale Ale", "Gose", "Quadrupel", "Tripel", "Lambic", "Rye Amber", "Kolsch",
			},
			"min_abv": 0.0,
			"max_abv": 15.0,
		},
	}

	err := ts.controller.GetAppConfig(ctx)

	ts.NoError(err)
	ts.Equal(http.StatusOK, rec.Code)

	var res map[string]interface{}
	err = json.Unmarshal([]byte(rec.Body.String()), &res)
	ts.NoError(err)
	ts.Equal(output, res)
}

func (ts *controllerTestSuite) buildContextAndRecorder(httpRequest *http.Request, userUserCase user.UseCase) (customCtx controller.CustomContext, rec *httptest.ResponseRecorder) {
	// echo.context가 아니라, controller.CustomContext 반환해야하는 것 유의
	rec = httptest.NewRecorder()
	ctx := ts.engine.NewContext(httpRequest, rec)
	customCtx = server.CustomContext{ctx, userUserCase}
	return
}

func (ts *controllerTestSuite) newHTTPRequest(method string, url string, params *url.Values, header *http.Header) http.Request {
	var httpRequest *http.Request
	var err error

	if method == http.MethodGet {
		httpRequest, err = http.NewRequest(method, url, nil)
		ts.NoError(err)
		ts.NotNil(httpRequest)
	} else {
		if params == nil {
			log.Fatalf("post http request needs body")
		}
		encodedParam := params.Encode()
		httpRequest, err = http.NewRequest(method, url, bytes.NewBufferString(encodedParam))
		ts.NoError(err)
		ts.NotNil(httpRequest)
	}

	if header != nil {
		httpRequest.Header = *header
	}
	return *httpRequest
}

type mockBeerUseCase struct {
	mock.Mock
}

func (u *mockBeerUseCase) AddBeer(beer beer.Beer) error {
	ret := u.Called(beer)
	return ret.Error(0)
}

func (u *mockBeerUseCase) GetBeers(args beer.BeerQueryArgs) ([]beer.Beer, error) {
	ret := u.Called(args)
	return ret.Get(0).([]beer.Beer), ret.Error(1)
}

func (u *mockBeerUseCase) GetBeer(beerID int64) (*beer.Beer, error) {
	ret := u.Called(beerID)
	return ret.Get(0).(*beer.Beer), ret.Error(1)
}

func (u *mockBeerUseCase) AddReview(review beer.Review) error {
	ret := u.Called(review)
	return ret.Error(0)
}

func (u *mockBeerUseCase) GetReviews(beerID int64) ([]beer.Review, error) {
	ret := u.Called(beerID)
	return ret.Get(0).([]beer.Review), ret.Error(1)
}

func (u *mockBeerUseCase) GetReviewsByUserID(userID int64) ([]beer.Review, error) {
	ret := u.Called(userID)
	return ret.Get(0).([]beer.Review), ret.Error(1)
}

func (u *mockBeerUseCase) GetReviewByBeerIDAndUserID(beerID int64, userID int64) (*beer.Review, error) {
	ret := u.Called(beerID, userID)
	return ret.Get(0).(*beer.Review), ret.Error(1)
}

func (u *mockBeerUseCase) GetRelatedBeers(beerID int64) (*beer.RelatedBeers, error) {
	ret := u.Called(beerID)
	return ret.Get(0).(*beer.RelatedBeers), ret.Error(1)
}

type mockUserUseCase struct {
	mock.Mock
}

func (u *mockUserUseCase) CreateUser(user user.User) error {
	ret := u.Called(user)
	return ret.Error(0)
}

func (u *mockUserUseCase) GetToken(code string) (*user.Token, error) {
	ret := u.Called(code)
	return ret.Get(0).(*user.Token), ret.Error(1)
}

func (u *mockUserUseCase) GetUser(accessToken string) (*user.User, error) {
	ret := u.Called(accessToken)
	return ret.Get(0).(*user.User), ret.Error(1)
}

func (u *mockUserUseCase) GetUserByID(userID int64) (*user.User, error) {
	ret := u.Called(userID)
	return ret.Get(0).(*user.User), ret.Error(1)
}

func (u *mockUserUseCase) GetUserByExternalID(externalID string) (*user.User, error) {
	ret := u.Called(externalID)
	return ret.Get(0).(*user.User), ret.Error(1)
}

func (u *mockUserUseCase) UpdateNickName(userID int64, nickName string) error {
	ret := u.Called(userID, nickName)
	return ret.Error(0)
}
