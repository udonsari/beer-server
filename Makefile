defualt: up

help:
	@echo 'Managements commands'

build:
	docker-compose build

up:
	docker-compose up

migrate:
	docker-compose run app bash -c './wait-for-it.sh mysqldb:3306 -- /bin/migration/table'

seed:
	docker-compose run app bash -c './wait-for-it.sh mysqldb:3306 -- /bin/migration/seed'

test-build:
	docker-compose -f docker-compose.test.yml build

test:
	docker-compose -f docker-compose.test.yml run test
