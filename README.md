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
    <details>
    <summary>Response Example</summary>
    <p>

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


---
### TODO
* CI / CD
* 코드 내부에 TODO 달아 놓은 것들
* 라이센스 추가
* REST API Convention 적용 점검
* Logger 사용
* Documentation
* Graceful Shutdown 처리 (서버, 외부 Dependency ...)
* `Develop ?` 리뷰 삭제 ?
* `Develop ?` 평균 별점 구간도 쿼리 ? 
* `Develop ?` 스타일 세분화 필요 ? ex) 대분류 에일, 중분류 IPA, 소분류 NEIPA
* `Develop ?` 소셜 로그인 연동 확장 ? (Factory 패턴. Naver, Google - external ID는 hashing 해서 provider별 prefix 달기, Model External ID 유니크 걸기)
* `Develop ?` 켜뮤니티 게시판 ?
* `Develop ?` 맥주 취향 추천 ?
* `Develop ?` 리뷰 달면 경헙치 -> 계급 올리는 개념 ?
* `중요` 토큰 Refresh, 만료 Client, Server 누가 처리하는지 알아보고 처리하기 (+로그아웃)
* `중요` 로그인 Token 자체를 Client에서 받게 하기. 서버는 Token 그냥 받고 (필요 없는 로직 지우기 - 근데 웹프론트에서는 필요할 것 같은데)

* `중요` RDS 분리
* `중요` 비슷한 맥주에 해당 맥주가 보인다.
* `중요` 이미지 가운데 정렬
* `중요` 실제 DB에 있는 향 리스트 내려주기

* `중요` 게시글 CRUD, 게시글 디테일, 댓글
* `중요` 와인그래프. App 찾아보기. 대댓글.
* `중요` 근처 맥주 바틀샵 찾기.

---
### On Going
* 미정님한테 맥주 이미지 수정 부탁 -> 완료시 반영하여 서버 업데이트
* 전반적으로 Validation 다듬기
    * DB 자체에 Name Unique 등
* Not fsound일 때 해당 객체 Nil 예외 처리 (ex. if err != nil 거르고, 바로 포인터 Dereference하지 않고 nil 체크)
* Test 구현 ( + Integration Test)

---

### Done
* `Done` 즐겨찾기 (추가 및 제거, 자기가 즐겨찾기 한 맥주 보기, 맥주 내려줄 때 자기가 즐겨찾기 했는지 체크)
* `Done` 10자로 닉네임 제한을 걸면 좋겠다
* `Done` 쿼리시 공백이 오면 어떻게 되지 ? -> 잘 됨
* `Done` 맥주 사이즈 정하고 (360 * 260), 그에 맞게 데이터 가공 하기 (이거 미정님 가이드 받고 배워두기. 이미지 사이즈에 맞게 재가공 하는 것, 누끼 따는 것 ... 그 외 미정님한테 도움될만한거 묻기 ) + 데이터 몇개 까지 (100개) ?
* `Done` Error 정의 및 대응되는 Status Code 사용 (ex. Auth Error)
* `Done` 마지막 페이지면, next_cursor null 내려주기
