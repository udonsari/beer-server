defualt: up

help:
	@echo 'Managements commands'

build:
	docker-compose build

up:
	docker-compose up

test-build:
	docker-compose -f docker-compose.test.yml build

test:
	docker-compose -f docker-compose.test.yml run test
