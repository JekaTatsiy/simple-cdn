build:
	cd center && GOOS=linux go build -o main main.go

up:	build
	docker compose build
	docker compose up

down: 
	docker compose down
