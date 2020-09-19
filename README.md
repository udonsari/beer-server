## Beer Server
---
* Make Command
    * `make build` : 서버용 Docker Image를 빌드 합니다
    * `make migrate-up` : 서버를 위한 MySQL Table을 Migration 합니다
    * `make migrate-down` : 서버를 위한 MySQL Table을 지웁니다
    * `make seed` : 서버를 위한 Beer Test Data를 넣습니다
    * `make up` : 서버를 실행합니다. 이후 `localhost:8081`로 접근 가능합니다 
    * `make test` : 코드 테스트를 실행 합니다
---
* API 예시
    * BeerList : 
        * command :`curl --location --request GET 'http://localhost:8081/api/beers?min_abv=5&max_abv=6&country=korea&beer_style=ipa&aroma=grape&beer_style=stout'`
            * aroma key, county, beer_style 등의 key를 중복 사용해서 해당 조건을 or로 걸 수 있습니다. 위 url에서는 beer_style이 중복 사용
        * Response Example
        ```json
        {
            "result": {
                "Beers": [
                    {
                        "id": 1,
                        "name": "Wonder Pale Ale",
                        "brewery": "CraftBros",
                        "abv": 5.7,
                        "country": "korea",
                        "beer_style": "ipa",
                        "aroma": [
                            "grape",
                            "apple"
                        ],
                        "image_url": [
                            "www.test_image_url.com"
                        ],
                        "comments": [
                            {
                                "beer_id": 1,
                                "content": "TEST_Comment_1",
                                "user_id": 1
                            },
                            {
                                "beer_id": 1,
                                "content": "TEST_Comment_2",
                                "user_id": 2
                            },
                            {
                                "beer_id": 1,
                                "content": "TEST_Comment_3",
                                "user_id": 3
                            }
                        ],
                        "rate_avg": 2.86,
                        "rate_owner": {
                            "BeerID": 1,
                            "Ratio": 2.364885155382621,
                            "UserID": 1
                        }
                    },
                    {
                        "id": 2,
                        "name": "Super Pale Ale",
                        "brewery": "CraftBros",
                        "abv": 6.3,
                        "country": "korea",
                        "beer_style": "ipa",
                        "aroma": [
                            "orange",
                            "apple"
                        ],
                        "image_url": [
                            "www.test_image_url.com"
                        ],
                        "comments": [
                            {
                                "beer_id": 2,
                                "content": "TEST_Comment_1",
                                "user_id": 1
                            },
                            {
                                "beer_id": 2,
                                "content": "TEST_Comment_2",
                                "user_id": 2
                            },
                            {
                                "beer_id": 2,
                                "content": "TEST_Comment_3",
                                "user_id": 3
                            }
                        ],
                        "rate_avg": 3.56,
                        "rate_owner": {
                            "BeerID": 2,
                            "Ratio": 2.655957595533158,
                            "UserID": 1
                        }
                    }
                ]
            }
        }
        ```
    * BeerDetail
        * command :`curl --location --request GET 'http://localhost:8081/api/beers?beer_id=5'`
            * beer_id를 인자로 줍니다.
        * Response Example
        ```json
        {
            "result": {
                "beer": {
                    "id": 101,
                    "name": "TEST_NAME_6026460253205598777",
                    "brewery": "TEST_BREWAERY_596",
                    "abv": 9.25,
                    "country": "TEST_COUNTRY_10",
                    "beer_style": "TEST_STYLE_13",
                    "aroma": [
                        "TEST_AROMA_8",
                        "TEST_AROMA_2",
                        "TEST_AROMA_8"
                    ],
                    "image_url": [
                        "http:naver.com",
                        "http:naver.com",
                        "http:naver.com",
                        "http:naver.com",
                        "http:naver.com"
                    ],
                    "comments": null,
                    "rate_avg": 0
                },
                "related_beers": {
                    "aroma_related": [
                        {
                            "id": 175,
                            "name": "TEST_NAME_966162726085428843",
                            "brewery": "TEST_BREWAERY_211",
                            "abv": 1.82,
                            "country": "TEST_COUNTRY_16",
                            "beer_style": "TEST_STYLE_9",
                            "aroma": [
                                "TEST_AROMA_2",
                                "TEST_AROMA_4",
                                "TEST_AROMA_5"
                            ],
                            "rate_avg": 0
                        },
                        {
                            "id": 125,
                            "name": "TEST_NAME_8659007057868230986",
                            "brewery": "TEST_BREWAERY_274",
                            "abv": 0.38,
                            "country": "TEST_COUNTRY_2",
                            "beer_style": "TEST_STYLE_4",
                            "aroma": [
                                "TEST_AROMA_2",
                                "TEST_AROMA_7",
                                "TEST_AROMA_8"
                            ],
                            "rate_avg": 0
                        },
                        {
                            "id": 172,
                            "name": "TEST_NAME_179496080759982046",
                            "brewery": "TEST_BREWAERY_570",
                            "abv": 5.34,
                            "country": "TEST_COUNTRY_20",
                            "beer_style": "TEST_STYLE_10",
                            "aroma": [
                                "TEST_AROMA_2",
                                "TEST_AROMA_1",
                                "TEST_AROMA_9"
                            ],
                            "rate_avg": 0
                        }
                    ],
                    "style_related": [
                        {
                            "id": 197,
                            "name": "TEST_NAME_4068733518307920492",
                            "brewery": "TEST_BREWAERY_988",
                            "abv": 4.25,
                            "country": "TEST_COUNTRY_21",
                            "beer_style": "TEST_STYLE_13",
                            "aroma": [
                                "TEST_AROMA_5",
                                "TEST_AROMA_8",
                                "TEST_AROMA_9"
                            ],
                            "rate_avg": 0
                        },
                        {
                            "id": 107,
                            "name": "TEST_NAME_1891709482230556383",
                            "brewery": "TEST_BREWAERY_289",
                            "abv": 8.45,
                            "country": "TEST_COUNTRY_21",
                            "beer_style": "TEST_STYLE_13",
                            "aroma": [
                                "TEST_AROMA_1",
                                "TEST_AROMA_9",
                                "TEST_AROMA_2"
                            ],
                            "rate_avg": 0
                        },
                        {
                            "id": 160,
                            "name": "TEST_NAME_2181965767827069027",
                            "brewery": "TEST_BREWAERY_909",
                            "abv": 5.72,
                            "country": "TEST_COUNTRY_15",
                            "beer_style": "TEST_STYLE_13",
                            "aroma": [
                                "TEST_AROMA_4",
                                "TEST_AROMA_4",
                                "TEST_AROMA_3"
                            ],
                            "rate_avg": 0
                        }
                    ],
                    "randomly_related": [
                        {
                            "id": 142,
                            "name": "TEST_NAME_4273616174478889151",
                            "brewery": "TEST_BREWAERY_779",
                            "abv": 0.16,
                            "country": "TEST_COUNTRY_11",
                            "beer_style": "TEST_STYLE_9",
                            "aroma": [
                                "TEST_AROMA_4",
                                "TEST_AROMA_7",
                                "TEST_AROMA_3"
                            ],
                            "rate_avg": 0
                        },
                        {
                            "id": 101,
                            "name": "TEST_NAME_6026460253205598777",
                            "brewery": "TEST_BREWAERY_596",
                            "abv": 9.25,
                            "country": "TEST_COUNTRY_10",
                            "beer_style": "TEST_STYLE_13",
                            "aroma": [
                                "TEST_AROMA_8",
                                "TEST_AROMA_2",
                                "TEST_AROMA_8"
                            ],
                            "rate_avg": 0
                        },
                        {
                            "id": 167,
                            "name": "TEST_NAME_3861280272082108960",
                            "brewery": "TEST_BREWAERY_104",
                            "abv": 1.64,
                            "country": "TEST_COUNTRY_20",
                            "beer_style": "TEST_STYLE_15",
                            "aroma": [
                                "TEST_AROMA_6",
                                "TEST_AROMA_7",
                                "TEST_AROMA_7"
                            ],
                            "rate_avg": 0
                        }
                    ]
                }
            }
        }        
        ```
    * SignIn (Kakao Only)
        * command :`curl --location --request GET 'http://localhost:8081/api/kakao/signin'` 
            * `api/token`로 Redirect되어 Access Token을 내려줍니다
            * 해당 토큰을 Header에 `Authorization`라는 Key의 Value로 담아 보내면 이후, 자신의 사용자 정보나 자신이 맥주에 매긴 Rate 등을 확인할 수 있습니다. Rate을 매기고, Comment를 달려면 마찬가지로 토큰을 설정해야합니다.
            * 로그인 연동을 테스트할 시, [연동 참조]라고 검색해서, 설명을 따라 주세요
        * Response Example
        ```json
        {
            "access_token": "ABC"
        }
        ```
    * UserDetail (계정 정보 뿌려주기, 추후 회원 탈퇴 만들기)
        * command :`curl --location --request GET 'http://localhost:8081/api/kakao/signin'` 
            * SignIn을 통해 얻은 Access Token을 Header에 담아 API 호출해야합니다
        * Response Example
        ```json
        {
            "ID": 1,
            "ExternalID": "0",
            "NickName": "",
            "ProfileImage": "",
            "ThumbnailImage": ""
        }        
        ```
    * TODO
        * BeerComment Post (SignIn 필요, 별점도 같이 [float], 글자 제한, 사용자랑 연관)
        * BeerComment List (개별 Beer ID를 인자로 옮)
            * 우선은 Comment와 Rate을 따로 뺌
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
* `Done` User API 내려주는 것 Snake Case 적용
* `Done` 카카오 Nickname이랑, 우리 Nickname 구분해야하는지 고민. (지금은 Kakao 닉네임 그대로 씀)
    * `Done` 사용자가 처음에 로그인하면 임의의 난수 Nickname 내려주고 (가마우지1379) 이후에 그걸 변경할 수 있도록 API 뚫어서 사용할 수 있게함 (안겹치게 Valdiation 유의)
* `Done` Comment, Rate Review로묶어서 처리 (유저당 맥주 마다 1번씩 처리하고, Upsert로 API)

* 맥주 Sorting도 해서 내려주는거 열기 (Comment 많은 순, Rate 높은 순)
* Pagination (다음 Page 호출하는 부분 처리, 애초에 Pagination 더 찾아보기)

* Comment, Rate 내려줄 때 Reduced Beer 정보, User NickName 내려주기 (기존 ID들은 같이)
* 사용자별 작성한 Comment, Rate 내려주는 API 뚫기
