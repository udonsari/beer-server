## Beer Server

* Make Command
    * `make build` : 서버용 Docker Image를 빌드 합니다
    * `make migrate` : 서버를 위한 MySQL Table을 Migration 합니다
    * `make seed` : 서버를 위한 Beer Test Data를 넣습니다
    * `make up` : 서버를 실행합니다. 이후 `localhost:8081`로 접근 가능합니다 

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
    * BeerDetail (Related Beers (향이 비슷, 맛이 비슷, 랜덤) 포함, 평균 별점 필요)
    * BeerComment Post (SignIn 필요, 별점도 같이 [float], 글자 제한, 사용자랑 연관)
    * BeerComment List (개별 Beer ID를 인자로 옮)
    * 우선은 Comment와 Rate을 따로 뺌

    * SignUp == SignIn
        * DB 실제 구성 필요
    * UserDetail (계정 정보 뿌려주기, 추후 회원 탈퇴 만들기)
        * DB 실제 구성 필요

* TODO
    * CI / CD
    * 코드 내부에 TODO 달아 놓은 것들
        * 우선순위 별로 별 달아놓음
    * 라이센스 추가
    * 전반적으로 Validation 다듬기
        * 한 맥주에 두 번 댓글 금지
        * DB 자체에 Name Unique 등
    * Error 정의 해서 사용
    * Test 구현
    * Logger 사용
    * 문서화
    * 에러 메시지 안내려가고 Internal Error로만 내려가는 것 처리

### 다음 PR은 DB 연결 및 확실한 내부 구현