build:
	docker build -t dns-test .

start:
	docker run --name=dns-dns -d  dns-test

up: build start
	docker exec -it dns-dns bash

kill:
	docker stop dns-dns
	docker rm dns-dns