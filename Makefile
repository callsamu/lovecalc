build: 
	docker-compose -f docker-compose.yml build

run: build
	docker-compose -f docker-compose.yml up

test: 
	docker-compose -f docker-compose.test.yml up --build
