## Beer Server

* Make Command
    * `make build` : 서버용 Docker Image를 빌드 합니다
    * `make up` : 서버를 실행합니다. 이후 `localhost:8081`로 접근 가능합니다

* API 예시
    * BeerList : 
        * command :`curl --location --request GET 'http://localhost:8081/api/beers?min_abv=5&max_abv=6&country=korea&beer_style=ipa&aroma=grape&beer_style=stout'`
            * aroma key, county, beer_style 등의 key를 중복 사용해서 해당 조건을 or로 걸 수 있습니다. 위 url에서는 beer_style이 중복 사용
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
            * 여기도 평점, 리뷰 개수 내려 보내기 (Sorting 위해). 일단 대략 완료
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
    * 라이센스 추가
    * 전반적으로 Validation 다듬기

### 다음 PR은 DB 연결 및 확실한 내부 구현