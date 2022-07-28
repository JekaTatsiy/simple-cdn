build:
	GOOS=linux go build -o main main.go

up:	
	docker compose build
	docker compose up

down: 
	docker compose down
