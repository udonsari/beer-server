## Beer Server

* Make Command
    * `make build` : 서버용 Docker Image를 빌드 합니다
    * `make up` : 서버를 실행합니다. 이후 localhost:8081로 접근 가능합니다

* API 예시
    * BeerList : `curl --location --request GET 'http://localhost:8081/api/beers?min_abv=5&max_abv=6&country=korea&beer_style=ipa&aroma=grape'`