## Beer Server

* Make Command
    * `make build` : 서버용 Docker Image를 빌드 합니다
    * `make up` : 서버를 실행합니다. 이후 `localhost:8081`로 접근 가능합니다

* API 예시
    * BeerList : `curl --location --request GET 'http://localhost:8081/api/beers?min_abv=5&max_abv=6&country=korea&beer_style=ipa&aroma=grape'`
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
    * CI / CD
    * 코드 내부에 TODO 달아 놓은 것들