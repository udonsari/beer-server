## Beer Server

* Make Command
    * `make build` : 서버용 Docker Image를 빌드 합니다
    * `make up` : 서버를 실행합니다. 이후 `localhost:8081`로 접근 가능합니다

* API 예시
    * BeerList : 
        * command :`curl --location --request GET 'http://localhost:8081/api/beers?min_abv=5&max_abv=6&country=korea&beer_style=ipa&aroma=grape'`
        * Response Example
        ```
        {
            "result": {
                "Beers": [
                    {
                        "name": "Wonder Pale Ale",
                        "brewary": "CraftBros",
                        "abv": 5.7,
                        "country": "korea",
                        "beer_style": "ipa",
                        "aroma": "grape"
                    },
                    {
                        "name": "Super Pale Ale",
                        "brewary": "CraftBros",
                        "abv": 6.3,
                        "country": "korea",
                        "beer_style": "ipa",
                        "aroma": "orange"
                    }
                ]
            }
        }
        ```
        * TODO
            * 여기도 평점, 리뷰 개수 내려 보내기 (Sorting 위해)
    * BeerDetail (Related Beer 포함, 평균 별점 필요)
    * BeerComment Post (SignIn 필요, 별점도 같이 [float], 글자 제한, 사용자랑 연관)
    * BeerComment List (개별 Beer ID를 인자로 옮)

    * SignUp == SignIn
        * DB 실제 구성 필요
    * UserDetail (계정 정보 뿌려주기, 추후 회원 탈퇴 만들기)
        * DB 실제 구성 필요

* TODO
    * CI / CD
    * 코드 내부에 TODO 달아 놓은 것들
    * 라이센스 추가