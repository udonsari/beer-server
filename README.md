# Beer Server


### Description
맥주에 대해 검색하고, 리뷰를 달고 서로의 리뷰를 살펴볼 수 있는 Server 입니다

--- 
### Make Command
* `make build` : 서버용 Docker Image를 빌드 합니다
* `make migrate-up` : 서버를 위한 MySQL Table을 Migration 합니다
* `make migrate-down` : 서버를 위한 MySQL Table을 지웁니다
* `make seed-fake` : 서버를 위한 Beer Test Data를 넣습니다
* `make seed-manual` : 서버를 위한 Beer Manual Data를 넣습니다
* `make seed` : 서버를 위한 Beer Data를 넣습니다 (현재 일부 Field는 의미 없는 값입니다)
* `make up` : 서버를 실행합니다. 이후 `localhost:8081`로 접근 가능합니다 
* `make test` : 코드 테스트를 실행 합니다
    <details>
    <summary>Current Test Coverage</summary>
    <p>

    ```bash
    go test ./... -coverprofile cover.out


    ?       github.com/UdonSari/beer-server/controller      [no test files]
    ok      github.com/UdonSari/beer-server/controller/beersvc      0.749s  coverage: 0.0% of statements [no tests to run]
    ?       github.com/UdonSari/beer-server/controller/beersvc/dto  [no test files]
    ok      github.com/UdonSari/beer-server/controller/usersvc      0.567s  coverage: 0.0% of statements [no tests to run]
    ?       github.com/UdonSari/beer-server/controller/usersvc/dto  [no test files]
    ok      github.com/UdonSari/beer-server/domain/beer     1.011s  coverage: 88.1% of statements
    ok      github.com/UdonSari/beer-server/domain/beer/repo        0.386s  coverage: 0.0% of statements [no tests to run]
    ok      github.com/UdonSari/beer-server/domain/user     0.375s  coverage: 0.0% of statements [no tests to run]
    ok      github.com/UdonSari/beer-server/domain/user/repo        0.993s  coverage: 0.0% of statements [no tests to run]
    ?       github.com/UdonSari/beer-server/main    [no test files]
    ok      github.com/UdonSari/beer-server/main/server     1.167s  coverage: 0.0% of statements [no tests to run]
    ?       github.com/UdonSari/beer-server/migration       [no test files]
    ?       github.com/UdonSari/beer-server/migration/commands      [no test files]
    ?       github.com/UdonSari/beer-server/util    [no test files]
    ```

    ```bash
    go tool cover -func cover.out


    github.com/UdonSari/beer-server/controller/beersvc/controller.go:25:    NewController                   0.0%
    github.com/UdonSari/beer-server/controller/beersvc/controller.go:39:    GetBeers                        0.0%
    github.com/UdonSari/beer-server/controller/beersvc/controller.go:109:   GetBeer                         0.0%
    github.com/UdonSari/beer-server/controller/beersvc/controller.go:186:   AddReview                       0.0%
    github.com/UdonSari/beer-server/controller/beersvc/controller.go:217:   GetReview                       0.0%
    github.com/UdonSari/beer-server/controller/beersvc/controller.go:251:   GetAppConfig                    0.0%
    github.com/UdonSari/beer-server/controller/beersvc/controller.go:261:   getDummyAppConfig               0.0%
    github.com/UdonSari/beer-server/controller/usersvc/controller.go:21:    NewController                   0.0%
    github.com/UdonSari/beer-server/controller/usersvc/controller.go:33:    SignInKakao                     0.0%
    github.com/UdonSari/beer-server/controller/usersvc/controller.go:47:    GetToken                        0.0%
    github.com/UdonSari/beer-server/controller/usersvc/controller.go:64:    GetUser                         0.0%
    github.com/UdonSari/beer-server/controller/usersvc/controller.go:79:    UpdateNickName                  0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:29:        New                             0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:36:        AddBeer                         0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:43:        GetBeer                         0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:59:        GetBeers                        0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:140:       UpdateBeerRateAvg               0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:145:       AddReview                       0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:171:       GetReviews                      0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:190:       GetReviewCount                  0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:196:       GetReviewByBeerIDAndUserID      0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/beerrepo.go:211:       GetReviewsByUserID              0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/mapper.go:12:          mapDBReviewToReview             0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/mapper.go:23:          mapReviewToDBReview             0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/mapper.go:32:          mapBeerToDBBeer                 0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/mapper.go:46:          mapDBBeerToBeer                 0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/mapper.go:62:          splitAndGetArray                0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/mapper.go:67:          splitAndGetString               0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/model.go:29:           TableName                       0.0%
    github.com/UdonSari/beer-server/domain/beer/repo/model.go:42:           TableName                       0.0%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:26:              NewUseCase                      100.0%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:33:              AddBeer                         100.0%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:37:              GetBeers                        100.0%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:41:              GetBeer                         100.0%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:45:              AddReview                       88.9%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:76:              GetReviews                      100.0%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:80:              GetReviewsByUserID              100.0%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:84:              GetReviewByBeerIDAndUserID      100.0%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:88:              GetRelatedBeers                 100.0%
    github.com/UdonSari/beer-server/domain/beer/usecase.go:123:             getRelatedBeersWithQueryArgs    100.0%
    github.com/UdonSari/beer-server/domain/beer/valueobject.go:45:          IsValidSortBy                   0.0%
    github.com/UdonSari/beer-server/domain/user/mapper.go:14:               NewMapper                       0.0%
    github.com/UdonSari/beer-server/domain/user/mapper.go:18:               MapKakaoUserToUser              0.0%
    github.com/UdonSari/beer-server/domain/user/mapper.go:30:               getRandomNickName               0.0%
    github.com/UdonSari/beer-server/domain/user/repo/mapper.go:7:           mapDBUserToUser                 0.0%
    github.com/UdonSari/beer-server/domain/user/repo/mapper.go:19:          mapUserToDBUser                 0.0%
    github.com/UdonSari/beer-server/domain/user/repo/model.go:16:           TableName                       0.0%
    github.com/UdonSari/beer-server/domain/user/repo/userrepo.go:20:        New                             0.0%
    github.com/UdonSari/beer-server/domain/user/repo/userrepo.go:26:        GetUserByExternalID             0.0%
    github.com/UdonSari/beer-server/domain/user/repo/userrepo.go:43:        GetUserByID                     0.0%
    github.com/UdonSari/beer-server/domain/user/repo/userrepo.go:58:        CreateUser                      0.0%
    github.com/UdonSari/beer-server/domain/user/repo/userrepo.go:69:        UpdateNickName                  0.0%
    github.com/UdonSari/beer-server/domain/user/usecase.go:37:              NewUseCase                      0.0%
    github.com/UdonSari/beer-server/domain/user/usecase.go:45:              CreateUser                      0.0%
    github.com/UdonSari/beer-server/domain/user/usecase.go:49:              GetToken                        0.0%
    github.com/UdonSari/beer-server/domain/user/usecase.go:82:              GetUser                         0.0%
    github.com/UdonSari/beer-server/domain/user/usecase.go:145:             GetUserByID                     0.0%
    github.com/UdonSari/beer-server/domain/user/usecase.go:149:             GetUserByExternalID             0.0%
    github.com/UdonSari/beer-server/domain/user/usecase.go:153:             UpdateNickName                  0.0%
    github.com/UdonSari/beer-server/main/server/customcontext.go:15:        User                            0.0%
    github.com/UdonSari/beer-server/main/server/customcontext.go:28:        UserMust                        0.0%
    github.com/UdonSari/beer-server/main/server/dependency.go:25:           NewDependency                   0.0%
    github.com/UdonSari/beer-server/main/server/dependency.go:29:           MysqlDB                         0.0%
    github.com/UdonSari/beer-server/main/server/dependency.go:48:           BeerCacheDuration               0.0%
    github.com/UdonSari/beer-server/main/server/dependency.go:52:           Host                            0.0%
    github.com/UdonSari/beer-server/main/server/dependency.go:56:           PortStr                         0.0%
    github.com/UdonSari/beer-server/main/server/dependency.go:60:           PortInt                         0.0%
    github.com/UdonSari/beer-server/main/server/dependency.go:64:           ServerEnv                       0.0%
    github.com/UdonSari/beer-server/main/server/dependency.go:68:           getEnvOrExit                    0.0%
    github.com/UdonSari/beer-server/main/server/dependency.go:76:           getInt64Env                     0.0%
    github.com/UdonSari/beer-server/main/server/server.go:32:               Init                            0.0%
    github.com/UdonSari/beer-server/main/server/server.go:38:               Start                           0.0%
    github.com/UdonSari/beer-server/main/server/server.go:51:               engine                          0.0%
    github.com/UdonSari/beer-server/main/server/server.go:90:               registerRoute                   0.0%
    github.com/UdonSari/beer-server/main/server/server.go:95:               New                             0.0%
    total:                                                                  (statements)                    10.1%
    ```
    </p>
    </details>
---
### Stack
* Language : Golang
* DB : MySQL
* ORM : gORM 
* Etc : Docker, Makefile
* Deployment : 단순히 AWS EC2 인스턴스에서 Git Pull 이후 Docker 빌드 및 실행합니다. (추후 변경)

---
### API 예시
* Nit.
    * 정확한 Parameter, Body는 각 Controller의 DTO를 봐주시면 좋습니다
* AppConfig
    * command : `curl --location --request GET 'http://localhost:8081/api/app-config'`
        * 클라이언트는 이 API를 호출하여 맥주 Filter UI를 위한 Aroma, Country, Style List를 얻어 초기화 합니다
        * 현재 version이라는 header를 보낼 경우와, 보내지 않는 경우 다른 AppConfig를 내립니다. 버전이 2가지 밖에 없으므로 version에는 임의의 값을 넣어도 상관 없으나 추후 Semantic Versioning에 따른 분기를 작업할 예정입니다
    <details>
    <summary>Response Example</summary>
    <p> V1

    ```json
    {
        "result": {
            "aroma_list": [
                "malty",
                "bicuity",
                "caramel",
                "roast",
                "coffee",
                "burnt",
                "grass",
                "blueberry",
                "banana",
                "pineapple",
                "apricot",
                "pear",
                "apple",
                "peach",
                "mango",
                "lemon",
                "orange",
                "grapefruit",
                "vinegar",
                "nutty"
            ],
            "country_list": [
                "USA",
                "Begium",
                "Genmany",
                "Korea",
                "UK",
                "Czech",
                "France"
            ],
            "style_list": [
                "Porter",
                "Stout",
                "Pilsener",
                "Light Lager",
                "Scotch Ale",
                "Saison",
                "Pale Ale",
                "Brown Ale",
                "India Pale Ale",
                "Gose",
                "Quadrupel",
                "Tripel",
                "Lambic"
            ],
            "min_abv": 0,
            "max_abv": 15
        }
    }
    ```
    </p>
    <p> V2

    ```json
    {
        "result": {
            "aroma_list": [
                "Malty",
                "Caramel",
                "Roast",
                "Coffee",
                "Grass",
                "Banana",
                "Apple",
                "Peach",
                "Mango",
                "Orange",
                "Spicy",
                "Vinegar",
                "Nutty",
                "Pineapple",
                "Melon",
                "Blackberry",
                "Chocolate",
                "Cherry",
                "Lemon",
                "Passion Fruit",
                "Grapefruit"
            ],
            "country_list": [
                "USA",
                "Begium",
                "Genmany",
                "Korea",
                "UK",
                "Czech",
                "France"
            ],
            "style_list": [
                {
                    "big_name": "Ale",
                    "mid_categories": [
                        {
                            "description": "상면 발효 효모를 사용하여\n화려하고 풍부한 향이 나는 맥주",
                            "mid_name": "Ale",
                            "small_categories": [
                                "Ale",
                                "Abbey Ale",
                                "Amber Ale",
                                "American Pale Ale",
                                "Brown Belgian Strong Ale",
                                "Blonde Ale",
                                "Brown Ale",
                                "Saison",
                                "Golden Ale",
                                "Hop Ale",
                                "Irish Ale",
                                "Light Ale",
                                "Old Ale",
                                "Pale Ale",
                                "Quadrupel Ale",
                                "Red Ale",
                                "Sparkling Ale",
                                "Summer Ale",
                                "Trappist Ale",
                                "Tripel Ale",
                                "White Ale",
                                "Wheat Ale",
                                "Wit Ale",
                                "Barley Wine",
                                "Dubbel Ale",
                                "Dark Ale",
                                "Wild Ale",
                                "Pumpkin Ale"
                            ]
                        },
                        {
                            "description": "페일 에일에 다량의 홉을 넣은,\n홉의 쌉쌀한 향과 맛이 매력적인 맥주",
                            "mid_name": "IPA",
                            "small_categories": [
                                "IPA",
                                "American IPA",
                                "Black IPA",
                                "Belgian IPA",
                                "Double IPA",
                                "Hazy IPA",
                                "Imperial IPA",
                                "Rye IPA",
                                "Session IPA",
                                "Sour IPA",
                                "Smoothie IPA",
                                "Wheat IPA"
                            ]
                        },
                        {
                            "description": "로스팅된 맥아를 사용한 어두운 색상의 맥주로\n풍부한 바디감이 특징인 맥주",
                            "mid_name": "Dark Beer",
                            "small_categories": [
                                "Dark Beer",
                                "Porter",
                                "Stout",
                                "Baltic Porter",
                                "Bourbon County Stout",
                                "Imperial Porter",
                                "Imperial Stout",
                                "Irish Stout",
                                "Sweet Stout",
                                "Schwarz",
                                "Milk Stout"
                            ]
                        },
                        {
                            "description": "밀 맥아를 높은 비율로 사용한 맥주로\n부드럽고 달콤한 향이 특징인 맥주",
                            "mid_name": "Wheat Beer",
                            "small_categories": [
                                "Wheat Beer",
                                "Belgian White",
                                "Hefeweizen",
                                "Witbier",
                                "Weizen",
                                "Dunkel Weizen",
                                "Weisse"
                            ]
                        }
                    ]
                },
                {
                    "big_name": "Larger",
                    "mid_categories": [
                        {
                            "description": "하면 발효 효모를 사용하여\n가벼운 풍미와 시원한 청량감이 매력적인 맥주",
                            "mid_name": "Larger",
                            "small_categories": [
                                "Lager",
                                "Amber Lager",
                                "Dark Lager",
                                "Helles Lager",
                                "India Pale Lager",
                                "Pale Lager",
                                "Rauchbier",
                                "Kellerbier",
                                "Marzen",
                                "Dunkel"
                            ]
                        },
                        {
                            "description": "다양한 원료와 긴 발효기간을 거쳐\n풍부한 맛과 높은 도수를 가진 맥주",
                            "mid_name": "Bock",
                            "small_categories": [
                                "Bock",
                                "Weizen Bock",
                                "Double Bock",
                                "MaiBock"
                            ]
                        }
                    ]
                },
                {
                    "big_name": "Lambic",
                    "mid_categories": [
                        {
                            "description": "상큼한 맛이 매력적인 자연 발효 맥주",
                            "mid_name": "Lambic",
                            "small_categories": [
                                "Lambic",
                                "Gueuze"
                            ]
                        }
                    ]
                },
                {
                    "big_name": "etc",
                    "mid_categories": [
                        {
                            "description": "비어있다에서 다양한 맥주를 만나보세요",
                            "mid_name": "etc",
                            "small_categories": [
                                "Radler",
                                "Cider",
                                "Gose",
                                "Gluten Free",
                                "Kolsch",
                                "Low Alcohol",
                                "Ginger Beer"
                            ]
                        }
                    ]
                }
            ],
            "min_abv": 0,
            "max_abv": 15
        }
    }
    ```
    </p>
    </details>
* Beer List
    * command : `curl --location --request GET 'http://localhost:8081/api/beers?min_abv=5&max_abv=6&beer_style=TEST_STYLE_1&aroma=TEST_AROMA_4&cursor=10&max_count=100&sort_by=review_count_desc'`
        * min_abv, max_abv를 명시하여 알콜 도수 범위를 제한할 수 있습니다
        * aroma, county, beer_style 등의 key를 중복 사용해서 해당 조건을 or로 걸 수 있습니다. 위 url에서는 beer_style이 중복 사용
        * Pagination - cursor는 첫 페이지의 경우 0으로 주고, 그 다음 호출 부터는 이 API의 Response cursor값을 넣으면 됩니다
        * SortBy - sort_by는 아래 값을 인자로 넣어주시면 됩니다
            ```go
            const (
                SortByRateAvgAsc      = "rate_avg_asc"
                SortByRateAvgDesc     = "rate_avg_desc"
                SortByReviewCountAsc  = "review_count_asc"
                SortByReviewCountDesc = "review_count_desc"
            )
            ```
    <details>
    <summary>Response Example</summary>
    <p>

    ```json
    {
        "result": {
            "beers": [
                {
                    "id": 9,
                    "name": "유레카 서울",
                    "brewery": "더부스",
                    "abv": 6.5,
                    "country": "Korea",
                    "beer_style": "ETC",
                    "aroma": [
                        "Peach"
                    ],
                    "thumbnail_image": "http://127.0.0.1:8081/static/thebooth_eurekaseoul.png",
                    "rate_avg": 0,
                    "review_count": 0,
                    "favorite_flag": false
                },
                {
                    "id": 10,
                    "name": "LIFE IPA 마릴린먼로",
                    "brewery": "크래프트브로스",
                    "abv": 6.5,
                    "country": "Korea",
                    "beer_style": "New England IPA",
                    "aroma": [
                        "Orange",
                        "Pineapple"
                    ],
                    "thumbnail_image": "http://127.0.0.1:8081/static/craftbros_lifeipamarilynmonroe.png",
                    "rate_avg": 4.5,
                    "review_count": 1,
                    "favorite_flag": false
                },
                {
                    "id": 11,
                    "name": "LIFE IPA 체게바라",
                    "brewery": "크래프트브로스",
                    "abv": 6.5,
                    "country": "Korea",
                    "beer_style": "New England IPA",
                    "aroma": [
                        "Orange",
                        "Pineapple"
                    ],
                    "thumbnail_image": "http://127.0.0.1:8081/static/craftbros_lifeipacheguevara.png",
                    "rate_avg": 0,
                    "review_count": 0,
                    "favorite_flag": true
                }
            ],
            "next_cursor": 11
        }
    }
    ```
    </p>
    </details>
* Beer Detail
    * command : `curl --location --request GET 'http://localhost:8081/api/beer?beer_id=30' --header 'Authorization: PodT6jlL4lAB6i_bT3lSyMnXguXbKMPeHjasdworDKcAAAF0h3rw4g'`
        * beer_id를 인자로 줍니다.
    <details>
    <summary>Response Example</summary>
    <p>

    ```json
    {
        "result": {
            "beer": {
                "id": 8,
                "name": "LGBTQ Smoothie IPA",
                "brewery": "더부스",
                "abv": 7.5,
                "country": "Korea",
                "beer_style": "India Pale Ale",
                "aroma": [
                    "Blackberry"
                ],
                "image_url": [
                    "http://127.0.0.1:8081/static/thebooth_lgbtqsmoothieipa.png"
                ],
                "thumbnail_image": "http://127.0.0.1:8081/static/thebooth_lgbtqsmoothieipa.png",
                "reviews": null,
                "rate_avg": 0,
                "review_count": 0,
                "favorite_flag": false
            },
            "related_beers": {
                "aroma_related": null,
                "style_related": [
                    {
                        "id": 1,
                        "name": "ㅋ IPA",
                        "brewery": "더부스",
                        "abv": 4.5,
                        "country": "Korea",
                        "beer_style": "India Pale Ale",
                        "aroma": [
                            "Peach",
                            "Mango"
                        ],
                        "thumbnail_image": "http://127.0.0.1:8081/static/thebooth_kieukipa.png",
                        "rate_avg": 0,
                        "review_count": 0,
                        "favorite_flag": false
                    },
                    {
                        "id": 5,
                        "name": "경리단 힙스터",
                        "brewery": "더부스",
                        "abv": 4.5,
                        "country": "Korea",
                        "beer_style": "India Pale Ale",
                        "aroma": [
                            "Orange",
                            "Grass"
                        ],
                        "thumbnail_image": "http://127.0.0.1:8081/static/thebooth_gyunglidanhipster.png",
                        "rate_avg": 0,
                        "review_count": 0,
                        "favorite_flag": false
                    }
                ],
                "randomly_related": [
                    {
                        "id": 3,
                        "name": "윗 마이 엑스",
                        "brewery": "더부스",
                        "abv": 5.5,
                        "country": "Korea",
                        "beer_style": "Witbier",
                        "aroma": [
                            "Orange",
                            "Spicy",
                            "Grass"
                        ],
                        "thumbnail_image": "http://127.0.0.1:8081/static/thebooth_witmyex.png",
                        "rate_avg": 0,
                        "review_count": 0,
                        "favorite_flag": false
                    },
                    {
                        "id": 4,
                        "name": "국민 IPA",
                        "brewery": "더부스",
                        "abv": 6.5,
                        "country": "Korea",
                        "beer_style": "India Pale Ale",
                        "aroma": [
                            "Orange",
                            "Pineapple",
                            "Melon"
                        ],
                        "thumbnail_image": "http://127.0.0.1:8081/static/thebooth_kookminipa.png",
                        "rate_avg": 0,
                        "review_count": 0,
                        "favorite_flag": false
                    },
                    {
                        "id": 1,
                        "name": "ㅋ IPA",
                        "brewery": "더부스",
                        "abv": 4.5,
                        "country": "Korea",
                        "beer_style": "India Pale Ale",
                        "aroma": [
                            "Peach",
                            "Mango"
                        ],
                        "thumbnail_image": "http://127.0.0.1:8081/static/thebooth_kieukipa.png",
                        "rate_avg": 0,
                        "review_count": 0,
                        "favorite_flag": false
                    }
                ]
            }
        }
    }
    ```
    </p>
    </details>
* Get RandomBeers
    * command : `curl --location --request GET 'http://localhost:8081/api/random-beers'`
        * Parameter를 가지지 않으며 Return 형태는 BeerList API와 같습니다. 다만, cursor는 매번 nil로 내려갑니다.
* Sign In (Kakao Only)
    * command : `curl --location --request GET 'http://localhost:8081/api/kakao/signin'` 
        * `api/token`로 Redirect되어 Access Token을 내려줍니다
        * 해당 토큰을 Header에 `Authorization`라는 Key의 Value로 담아 보내면 이후, 자신의 사용자 정보나 자신이 맥주에 매긴 Rate 등을 확인할 수 있습니다. Rate을 매기고, Comment를 달려면 마찬가지로 토큰을 설정해야합니다.
        * 로그인 연동을 테스트할 시, [연동 참조]라고 검색해서, 설명을 따라 주세요
    <details>
    <summary>Response Example</summary>
    <p>

    ```json
    {
        "access_token": "ABC"
    }
    ```
    </p>
    </details>
* User Detail (For Logined User)
    * command : `curl --location --request GET 'http://localhost:8081/api/user' --header 'Authorization: PodT6jlL4lAB6i_bT3lSyMnXguXbKMPeHjasdworDKcAAAF0h3rw4g'` 
        * SignIn을 통해 얻은 Access Token을 Header에 담아 API 호출해야합니다
    <details>
    <summary>Response Example</summary>
    <p>

    ```json
    {
        "result": {
            "id": 101,
            "external_id": "0",
            "nickname": "Crownbig-1976380737404962472",
            "profile_image": "",
            "thumbnail_image": ""
        }
    }
    ```
    </p>
    </details>
* Update User Nickname (For Logined User)
    * command : `curl --location --request POST 'http://localhost:8081/api/user/update' --header 'Authorization: PodT6jlL4lAB6i_bT3lSyMnXguXbKMPeHjasdworDKcAAAF0h3rw4g' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'nickname=Crownbig-1976380737404962472'`
    * Resonse : None
* Add Review (For Logined User)
    * command : `curl --location --request POST 'http://localhost:8081/api/review' --header 'Authorization: TWcoWfUC5WoyHXPdlc37kMtAZww5gNWJNQAXuQo9c5oAAAF0hoDyDg' --form 'beer_id=30' --form 'ratio=4.2' --form 'content=JUST_3'`
    * Response : None
* Get Review (For Logined User)
    * command : `curl --location --request GET 'http://localhost:8081/api/review' --header 'Authorization: TWcoWfUC5WoyHXPdlc37kMtAZww5gNWJNQAXuQo9c5oAAAF0hoDyDg'`
    <details>
    <summary>Response Example</summary>
    <p>

    ```json
    {
        "result": [
            {
                "beer": {
                    "id": 30,
                    "name": "TEST_NAME_4137880265740432633",
                    "brewery": "TEST_BREWAERY_86",
                    "abv": 2.85,
                    "country": "TEST_COUNTRY_0",
                    "beer_style": "TEST_STYLE_4",
                    "aroma": [
                        "TEST_AROMA_1",
                        "TEST_AROMA_2",
                        "TEST_AROMA_1"
                    ],
                    "thumbnail_image": "https://picsum.photos/320/480",
                    "rate_avg": 3.64
                },
                "content": "JUST_3",
                "ratio": 4.2,
                "user_id": 101,
                "nickname": "Crownbig-1976380737404962472"
            },
            {
                "beer": {
                    "id": 33,
                    "name": "TEST_NAME_520284185256194436",
                    "brewery": "TEST_BREWAERY_78",
                    "abv": 9.52,
                    "country": "TEST_COUNTRY_5",
                    "beer_style": "TEST_STYLE_1",
                    "aroma": [
                        "TEST_AROMA_3",
                        "TEST_AROMA_3",
                        "TEST_AROMA_3"
                    ],
                    "thumbnail_image": "https://picsum.photos/320/480",
                    "rate_avg": 3.73
                },
                "content": "JUST_4",
                "ratio": 4.2,
                "user_id": 101,
                "nickname": "Crownbig-1976380737404962472"
            }
        ]
    }
    ```
    </p>
    </details>
* Add Favorite
    * command : `curl --location --request POST 'http://localhost:8081/api/favorite' --header 'Authorization: TWcoWfUC5WoyHXPdlc37kMtAZww5gNWJNQAXuQo9c5oAAAF0hoDyDg' --form 'beer_id="10"'  --form 'flag="false"'`
    * Response : None
* Get Favorites
    * command : `curl --location --request GET 'http://localhost:8081/api/favorite' --header 'Authorization: TWcoWfUC5WoyHXPdlc37kMtAZww5gNWJNQAXuQo9c5oAAAF0hoDyDg'`
    <details>
    <summary>Response Example</summary>
    <p>

    ```json
    {
        "result": [
            {
                "beer": {
                    "id": 11,
                    "name": "LIFE IPA 체게바라",
                    "brewery": "크래프트브로스",
                    "abv": 6.5,
                    "country": "Korea",
                    "beer_style": "New England IPA",
                    "aroma": [
                        "Orange",
                        "Pineapple"
                    ],
                    "thumbnail_image": "http://127.0.0.1:8081/static/craftbros_lifeipacheguevara.png",
                    "rate_avg": 0,
                    "review_count": 0,
                    "favorite_flag": true
                },
                "user_id": 1,
                "beer_id": 11
            },
            {
                "beer": {
                    "id": 15,
                    "name": "헌치백 세션 IPA",
                    "brewery": "플레이그라운드",
                    "abv": 4,
                    "country": "Korea",
                    "beer_style": "India Pale Ale",
                    "aroma": [
                        "Orange"
                    ],
                    "thumbnail_image": "http://127.0.0.1:8081/static/playground_hunchback.png",
                    "rate_avg": 0,
                    "review_count": 0,
                    "favorite_flag": true
                },
                "user_id": 1,
                "beer_id": 15
            }
        ]
    }
    ```
    </p>
    </details>
* Add User Beer Config
    * command : `curl --location --request POST 'http://localhost:8081/api/user-beer-config' --header 'Authorization: TWcoWfUC5WoyHXPdlc37kMtAZww5gNWJNQAXuQo9c5oAAAF0hoDyDg' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'aroma=aroma1, aroma2' --data-urlencode 'style=style1, style2'`
    * Response : None
* Get User Beer Config
    * command : `curl --location --request GET 'http://localhost:8081/api/user-beer-config' --header 'Authorization: TWcoWfUC5WoyHXPdlc37kMtAZww5gNWJNQAXuQo9c5oAAAF0hoDyDg'`
    <details>
    <summary>Response Example</summary>
    <p>

    ```json
    {
        "result": {
            "Aroma": [
                "aroma1, aroma2"
            ],
            "Style": [
                "style1, style2"
            ]
        }
    }
    ```
    </p>
    </details>
* Get Popular Beers
    * command : `curl --location --request GET 'http://localhost:8081/api/popular-beers?start_date=2000-01-01%2000:00:00&end_date=2100-01-01%2000:00:00&limit=5' --header 'Authorization: TWcoWfUC5WoyHXPdlc37kMtAZww5gNWJNQAXuQo9c5oAAAF0hoDyDg'`
        * 이 API는 start_date, end_date로 명시되는 기간에서 가장 좋아요를 많이 받은 맥주리스트를 반환합니다.
        * 날짜 포맷은 `2100-01-01 00:00:00` 와 같은 포맷이어야합니다.
        * 만약 start_date, end_date를 포함하지 않고 API를 호출하면 기간은 현재 달로 설정됩니다.
    * Response : beer list api와 return format이 같으나, cursor 값은 의미 없습니다.


---
