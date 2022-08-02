build:
	cd center && GOOS=linux go build -o main main.go

up:	build
	docker compose build
	docker compose up

down: 
	docker compose down

cert:
	cd ./key && bash create_cert.sh site.com