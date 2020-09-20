## Beer Server

### Make Command
* `make build` : 서버용 Docker Image를 빌드 합니다
* `make migrate-up` : 서버를 위한 MySQL Table을 Migration 합니다
* `make migrate-down` : 서버를 위한 MySQL Table을 지웁니다
* `make seed` : 서버를 위한 Beer Test Data를 넣습니다
* `make up` : 서버를 실행합니다. 이후 `localhost:8081`로 접근 가능합니다 
* `make test` : 코드 테스트를 실행 합니다
---
### Stack
* Language : Golang
* DB : MySQL
* ORM : gORM 
* Etc : Docker, Makefile
---
### API 예시
* Nit.
    * 정확한 Parameter, Body는 각 Controller의 DTO를 봐주시면 좋습니다
* Beer List
    * command : `curl --location --request GET 'http://localhost:8081/api/beers?min_abv=5&max_abv=6&cursor=61&max_count=2'`
        * aroma, county, beer_style 등의 key를 중복 사용해서 해당 조건을 or로 걸 수 있습니다. 위 url에서는 beer_style이 중복 사용
        * Pagination - cursor는 첫 페이지의 경우 0으로 주고, 그 다음 호출 부터는 이 API의 Response cursor값을 넣으면 됩니다
    <details>
    <summary>Response Example</summary>
    <p>

    ```json
    {
        "result": {
            "beers": [
                {
                    "id": 79,
                    "name": "TEST_NAME_6359516315660856847",
                    "brewery": "TEST_BREWAERY_70",
                    "abv": 5.37,
                    "country": "TEST_COUNTRY_3",
                    "beer_style": "TEST_STYLE_4",
                    "aroma": [
                        "TEST_AROMA_3",
                        "TEST_AROMA_3",
                        "TEST_AROMA_3"
                    ],
                    "image_url": [
                        "https://picsum.photos/320/480",
                        "https://picsum.photos/320/480",
                        "https://picsum.photos/320/480",
                        "https://picsum.photos/320/480",
                        "https://picsum.photos/320/480"
                    ],
                    "thumbnail_image": "https://picsum.photos/320/480",
                    "reviews": [
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_7388582044909595896",
                            "ratio": 4.59,
                            "user_id": 19,
                            "nickname": "TEST_NICKNAME_2399578652595903116"
                        },
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_7036775995447857290",
                            "ratio": 3.8,
                            "user_id": 20,
                            "nickname": "TEST_NICKNAME_9184890450261584277"
                        },
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_2391156082219578938",
                            "ratio": 3.69,
                            "user_id": 25,
                            "nickname": "TEST_NICKNAME_1916903903388503463"
                        },
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_5924617277191323549",
                            "ratio": 3.19,
                            "user_id": 43,
                            "nickname": "TEST_NICKNAME_3210880343281673809"
                        },
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_6452731955647185282",
                            "ratio": 4.94,
                            "user_id": 45,
                            "nickname": "TEST_NICKNAME_9191961633212244318"
                        },
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_5801703947891575618",
                            "ratio": 2.23,
                            "user_id": 46,
                            "nickname": "TEST_NICKNAME_4628425314673383947"
                        },
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_8647855132242115017",
                            "ratio": 4.57,
                            "user_id": 48,
                            "nickname": "TEST_NICKNAME_4133041801356895198"
                        },
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_2298469925369900087",
                            "ratio": 4.39,
                            "user_id": 53,
                            "nickname": "TEST_NICKNAME_4300038458279069427"
                        },
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_4013933872477880882",
                            "ratio": 4.61,
                            "user_id": 68,
                            "nickname": "TEST_NICKNAME_542935270987147634"
                        },
                        {
                            "beer": {
                                "id": 79,
                                "name": "TEST_NAME_6359516315660856847",
                                "brewery": "TEST_BREWAERY_70",
                                "abv": 5.37,
                                "country": "TEST_COUNTRY_3",
                                "beer_style": "TEST_STYLE_4",
                                "aroma": [
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3",
                                    "TEST_AROMA_3"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.9
                            },
                            "content": "TEST_CONTENT_6662959013763429116",
                            "ratio": 2.97,
                            "user_id": 97,
                            "nickname": "TEST_NICKNAME_7599357322839615581"
                        }
                    ],
                    "rate_avg": 3.9
                },
                {
                    "id": 95,
                    "name": "TEST_NAME_723102351307689977",
                    "brewery": "TEST_BREWAERY_5",
                    "abv": 5.1,
                    "country": "TEST_COUNTRY_7",
                    "beer_style": "TEST_STYLE_2",
                    "aroma": [
                        "TEST_AROMA_1",
                        "TEST_AROMA_4",
                        "TEST_AROMA_1"
                    ],
                    "image_url": [
                        "https://picsum.photos/320/480",
                        "https://picsum.photos/320/480",
                        "https://picsum.photos/320/480",
                        "https://picsum.photos/320/480",
                        "https://picsum.photos/320/480"
                    ],
                    "thumbnail_image": "https://picsum.photos/320/480",
                    "reviews": [
                        {
                            "beer": {
                                "id": 95,
                                "name": "TEST_NAME_723102351307689977",
                                "brewery": "TEST_BREWAERY_5",
                                "abv": 5.1,
                                "country": "TEST_COUNTRY_7",
                                "beer_style": "TEST_STYLE_2",
                                "aroma": [
                                    "TEST_AROMA_1",
                                    "TEST_AROMA_4",
                                    "TEST_AROMA_1"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.74
                            },
                            "content": "TEST_CONTENT_5961655009417013599",
                            "ratio": 3.4,
                            "user_id": 22,
                            "nickname": "TEST_NICKNAME_4615009674455111748"
                        },
                        {
                            "beer": {
                                "id": 95,
                                "name": "TEST_NAME_723102351307689977",
                                "brewery": "TEST_BREWAERY_5",
                                "abv": 5.1,
                                "country": "TEST_COUNTRY_7",
                                "beer_style": "TEST_STYLE_2",
                                "aroma": [
                                    "TEST_AROMA_1",
                                    "TEST_AROMA_4",
                                    "TEST_AROMA_1"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.74
                            },
                            "content": "TEST_CONTENT_986019216809696119",
                            "ratio": 4.83,
                            "user_id": 23,
                            "nickname": "TEST_NICKNAME_5345315911668805533"
                        },
                        {
                            "beer": {
                                "id": 95,
                                "name": "TEST_NAME_723102351307689977",
                                "brewery": "TEST_BREWAERY_5",
                                "abv": 5.1,
                                "country": "TEST_COUNTRY_7",
                                "beer_style": "TEST_STYLE_2",
                                "aroma": [
                                    "TEST_AROMA_1",
                                    "TEST_AROMA_4",
                                    "TEST_AROMA_1"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.74
                            },
                            "content": "TEST_CONTENT_1913799607136721406",
                            "ratio": 4.63,
                            "user_id": 25,
                            "nickname": "TEST_NICKNAME_1916903903388503463"
                        },
                        {
                            "beer": {
                                "id": 95,
                                "name": "TEST_NAME_723102351307689977",
                                "brewery": "TEST_BREWAERY_5",
                                "abv": 5.1,
                                "country": "TEST_COUNTRY_7",
                                "beer_style": "TEST_STYLE_2",
                                "aroma": [
                                    "TEST_AROMA_1",
                                    "TEST_AROMA_4",
                                    "TEST_AROMA_1"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.74
                            },
                            "content": "TEST_CONTENT_8837695456469781027",
                            "ratio": 2.73,
                            "user_id": 29,
                            "nickname": "TEST_NICKNAME_2348336246068030497"
                        },
                        {
                            "beer": {
                                "id": 95,
                                "name": "TEST_NAME_723102351307689977",
                                "brewery": "TEST_BREWAERY_5",
                                "abv": 5.1,
                                "country": "TEST_COUNTRY_7",
                                "beer_style": "TEST_STYLE_2",
                                "aroma": [
                                    "TEST_AROMA_1",
                                    "TEST_AROMA_4",
                                    "TEST_AROMA_1"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.74
                            },
                            "content": "TEST_CONTENT_2143404436982914079",
                            "ratio": 3.59,
                            "user_id": 62,
                            "nickname": "TEST_NICKNAME_4459496347945300277"
                        },
                        {
                            "beer": {
                                "id": 95,
                                "name": "TEST_NAME_723102351307689977",
                                "brewery": "TEST_BREWAERY_5",
                                "abv": 5.1,
                                "country": "TEST_COUNTRY_7",
                                "beer_style": "TEST_STYLE_2",
                                "aroma": [
                                    "TEST_AROMA_1",
                                    "TEST_AROMA_4",
                                    "TEST_AROMA_1"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.74
                            },
                            "content": "TEST_CONTENT_1004243693761545598",
                            "ratio": 3.4,
                            "user_id": 67,
                            "nickname": "TEST_NICKNAME_7781229485883467616"
                        },
                        {
                            "beer": {
                                "id": 95,
                                "name": "TEST_NAME_723102351307689977",
                                "brewery": "TEST_BREWAERY_5",
                                "abv": 5.1,
                                "country": "TEST_COUNTRY_7",
                                "beer_style": "TEST_STYLE_2",
                                "aroma": [
                                    "TEST_AROMA_1",
                                    "TEST_AROMA_4",
                                    "TEST_AROMA_1"
                                ],
                                "thumbnail_image": "https://picsum.photos/320/480",
                                "rate_avg": 3.74
                            },
                            "content": "TEST_CONTENT_8085769321972188030",
                            "ratio": 3.62,
                            "user_id": 87,
                            "nickname": "TEST_NICKNAME_6235857566738269820"
                        }
                    ],
                    "rate_avg": 3.74
                }
            ],
            "next_cursor": 95
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
                "image_url": [
                    "https://picsum.photos/320/480",
                    "https://picsum.photos/320/480",
                    "https://picsum.photos/320/480",
                    "https://picsum.photos/320/480",
                    "https://picsum.photos/320/480"
                ],
                "thumbnail_image": "https://picsum.photos/320/480",
                "reviews": [
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
                        "content": "TEST_CONTENT_2555230713823100474",
                        "ratio": 2.96,
                        "user_id": 18,
                        "nickname": "TEST_NICKNAME_2865495138003791683"
                    },
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
                        "content": "TEST_CONTENT_2013357936790157503",
                        "ratio": 3.01,
                        "user_id": 20,
                        "nickname": "TEST_NICKNAME_9184890450261584277"
                    },
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
                        "content": "TEST_CONTENT_1358455889342800964",
                        "ratio": 3.06,
                        "user_id": 28,
                        "nickname": "TEST_NICKNAME_646140600132340307"
                    },
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
                        "content": "TEST_CONTENT_6022346283882884485",
                        "ratio": 3.12,
                        "user_id": 30,
                        "nickname": "TEST_NICKNAME_1667525891839284287"
                    },
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
                        "content": "TEST_CONTENT_2816295878650713538",
                        "ratio": 4.68,
                        "user_id": 33,
                        "nickname": "TEST_NICKNAME_2286410554287199904"
                    },
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
                        "content": "TEST_CONTENT_6329023503589439101",
                        "ratio": 4.51,
                        "user_id": 36,
                        "nickname": "TEST_NICKNAME_5632003992361113250"
                    },
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
                        "content": "TEST_CONTENT_1913796447148404678",
                        "ratio": 4.12,
                        "user_id": 44,
                        "nickname": "TEST_NICKNAME_5620865574498117321"
                    },
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
                        "content": "TEST_CONTENT_5578483510544644879",
                        "ratio": 3.54,
                        "user_id": 47,
                        "nickname": "TEST_NICKNAME_8641665568174466744"
                    },
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
                        "content": "TEST_CONTENT_4292709364036596106",
                        "ratio": 4.45,
                        "user_id": 63,
                        "nickname": "TEST_NICKNAME_6244968064195571465"
                    },
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
                        "content": "TEST_CONTENT_8203359081787729076",
                        "ratio": 2.97,
                        "user_id": 65,
                        "nickname": "TEST_NICKNAME_4908180098745379294"
                    },
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
                        "content": "TEST_CONTENT_8483473889263299268",
                        "ratio": 3.04,
                        "user_id": 78,
                        "nickname": "TEST_NICKNAME_1744005986462343804"
                    },
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
                        "content": "TEST_CONTENT_2050597700810524849",
                        "ratio": 4.55,
                        "user_id": 86,
                        "nickname": "TEST_NICKNAME_2883257872911213149"
                    },
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
                        "content": "TEST_CONTENT_7976959338559980696",
                        "ratio": 3.35,
                        "user_id": 99,
                        "nickname": "TEST_NICKNAME_8219864395714989725"
                    },
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
                    }
                ],
                "rate_avg": 3.64,
                "review_owner": {
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
                }
            },
            "related_beers": {
                "aroma_related": [
                    {
                        "id": 57,
                        "name": "TEST_NAME_368154985306580155",
                        "brewery": "TEST_BREWAERY_64",
                        "abv": 8.81,
                        "country": "TEST_COUNTRY_7",
                        "beer_style": "TEST_STYLE_2",
                        "aroma": [
                            "TEST_AROMA_1",
                            "TEST_AROMA_2",
                            "TEST_AROMA_0"
                        ],
                        "thumbnail_image": "https://picsum.photos/320/480",
                        "rate_avg": 3.37
                    },
                    {
                        "id": 4,
                        "name": "TEST_NAME_2636780126543146206",
                        "brewery": "TEST_BREWAERY_86",
                        "abv": 2.27,
                        "country": "TEST_COUNTRY_3",
                        "beer_style": "TEST_STYLE_1",
                        "aroma": [
                            "TEST_AROMA_1",
                            "TEST_AROMA_2",
                            "TEST_AROMA_1"
                        ],
                        "thumbnail_image": "https://picsum.photos/320/480",
                        "rate_avg": 2.79
                    },
                    {
                        "id": 45,
                        "name": "TEST_NAME_1297920382251099222",
                        "brewery": "TEST_BREWAERY_41",
                        "abv": 8.48,
                        "country": "TEST_COUNTRY_5",
                        "beer_style": "TEST_STYLE_0",
                        "aroma": [
                            "TEST_AROMA_1",
                            "TEST_AROMA_0",
                            "TEST_AROMA_2"
                        ],
                        "thumbnail_image": "https://picsum.photos/320/480",
                        "rate_avg": 3.6
                    }
                ],
                "style_related": [
                    {
                        "id": 60,
                        "name": "TEST_NAME_6490111625982009499",
                        "brewery": "TEST_BREWAERY_71",
                        "abv": 3.03,
                        "country": "TEST_COUNTRY_2",
                        "beer_style": "TEST_STYLE_4",
                        "aroma": [
                            "TEST_AROMA_4",
                            "TEST_AROMA_1",
                            "TEST_AROMA_3"
                        ],
                        "thumbnail_image": "https://picsum.photos/320/480",
                        "rate_avg": 3.59
                    },
                    {
                        "id": 68,
                        "name": "TEST_NAME_4087782157814699375",
                        "brewery": "TEST_BREWAERY_52",
                        "abv": 4.81,
                        "country": "TEST_COUNTRY_4",
                        "beer_style": "TEST_STYLE_4",
                        "aroma": [
                            "TEST_AROMA_0",
                            "TEST_AROMA_1",
                            "TEST_AROMA_1"
                        ],
                        "thumbnail_image": "https://picsum.photos/320/480",
                        "rate_avg": 3.7
                    },
                    {
                        "id": 16,
                        "name": "TEST_NAME_3684883192478116607",
                        "brewery": "TEST_BREWAERY_22",
                        "abv": 7.83,
                        "country": "TEST_COUNTRY_0",
                        "beer_style": "TEST_STYLE_4",
                        "aroma": [
                            "TEST_AROMA_4",
                            "TEST_AROMA_3",
                            "TEST_AROMA_1"
                        ],
                        "thumbnail_image": "https://picsum.photos/320/480",
                        "rate_avg": 3.24
                    }
                ],
                "randomly_related": [
                    {
                        "id": 51,
                        "name": "TEST_NAME_6629569320063026401",
                        "brewery": "TEST_BREWAERY_12",
                        "abv": 3.01,
                        "country": "TEST_COUNTRY_4",
                        "beer_style": "TEST_STYLE_2",
                        "aroma": [
                            "TEST_AROMA_4",
                            "TEST_AROMA_4",
                            "TEST_AROMA_1"
                        ],
                        "thumbnail_image": "https://picsum.photos/320/480",
                        "rate_avg": 3.88
                    },
                    {
                        "id": 81,
                        "name": "TEST_NAME_6832561201158269111",
                        "brewery": "TEST_BREWAERY_71",
                        "abv": 1.72,
                        "country": "TEST_COUNTRY_5",
                        "beer_style": "TEST_STYLE_2",
                        "aroma": [
                            "TEST_AROMA_3",
                            "TEST_AROMA_2",
                            "TEST_AROMA_3"
                        ],
                        "thumbnail_image": "https://picsum.photos/320/480",
                        "rate_avg": 3.32
                    },
                    {
                        "id": 29,
                        "name": "TEST_NAME_5918515489820760331",
                        "brewery": "TEST_BREWAERY_78",
                        "abv": 2.08,
                        "country": "TEST_COUNTRY_0",
                        "beer_style": "TEST_STYLE_0",
                        "aroma": [
                            "TEST_AROMA_4",
                            "TEST_AROMA_4",
                            "TEST_AROMA_0"
                        ],
                        "thumbnail_image": "https://picsum.photos/320/480",
                        "rate_avg": 4.45
                    }
                ]
            }
        }
    }
    ```
    </p>
    </details>
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
* User Detail
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
* Update User Nickname
    * command : `curl --location --request POST 'http://localhost:8081/api/user/update' --header 'Authorization: PodT6jlL4lAB6i_bT3lSyMnXguXbKMPeHjasdworDKcAAAF0h3rw4g' --header 'Content-Type: application/x-www-form-urlencoded' --data-urlencode 'nickname=Crownbig-1976380737404962472'`
    * Resonse : None
* Add Review
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
---
* TODO
    * CI / CD
    * 코드 내부에 TODO 달아 놓은 것들
        * 우선순위 별로 별 달아놓음
    * 라이센스 추가
    * REST API Convention 적용 점검
    * Test 구현
    * Logger 사용
    * 문서화
    * Graceful Shutdown 처리 (서버, 외부 Dependency ...)
    * 리뷰 삭제 ? - 일단은 없다.
    * 평균 별점 구간도 쿼리 할 수 있어야겠네
    * Not found일 때 해당 객체 Nil 예외 처리 (ex. if err != nil 거르고, 바로 포인터 Dereference하지 않고 nil 체크)
    * `Develop` 스타일 세분화 필요. ex) 대분류 에일, 중분류 IPA, 소분류 NEIPA
    * `Develop` 소셜 로그인 연동 확장 (Factory 패턴. Naver, Google - external ID는 hashing 해서 provider별 prefix 달기, Model External ID 유니크 걸기)
    * `Develop` 켜뮤니티
    * `중요` Error 정의 및 대응되는 Status Code 사용 (ex. Auth Error)
    * `중요` 전반적으로 Validation 다듬기
        * 한 맥주에 두 번 댓글 금지
        * DB 자체에 Name Unique 등
    * `중요` 맥주 데이터 넣기 - 기한 27일
    * `중요` AWS 서버 띄우는 건 - 기한 27일
    * `중요` 토큰 Refresh, 만료 Client, Server 누가 처리하는지 알아보고 처리하기 (+로그아웃)
    * `중요` 로그인 Token 자체를 Client에서 받게 하기. 서버는 Token 그냥 받고 (필요 없는 로직 지우기 - 근데 웹프론트에서는 필요할 것 같은데)
---
* On Going
    * 맥주 Sorting도 해서 내려주는거 열기 (Comment 많은 순, RateAvg 높은 순)
        * Comment 많은 순은 쿼리를 Join해서 좀 만져야할 듯
    