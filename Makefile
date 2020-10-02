defualt: up

help:
	@echo 'Managements commands'

build:
	docker-compose build

up:
	docker-compose up

migrate-down:
	docker-compose run app bash -c './wait-for-it.sh mysqldb:3306 -- /bin/migration migrate-down'

migrate-up:
	docker-compose run app bash -c './wait-for-it.sh mysqldb:3306 -- /bin/migration migrate-up'

seed-fake:
	docker-compose run app bash -c './wait-for-it.sh mysqldb:3306 -- /bin/migration seed-fake'

seed:
	docker-compose run app bash -c './wait-for-it.sh mysqldb:3306 -- /bin/migration seed'

test-build:
	docker-compose -f docker-compose.test.yml build

test:
	docker-compose -f docker-compose.test.yml run test

# Unofficial
refresh:
	make build
	make migrate-down
	make migrate-up
	make seed

refresh-fake:
	make build
	make migrate-down
	make migrate-up
	make seed-fake
