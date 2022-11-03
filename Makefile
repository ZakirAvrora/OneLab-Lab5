build:
	docker build --rm -t crud-web .
	docker image prune --filter label=stage=builder -f
run:
	docker run --rm --name crud-web -p 8080:8080 crud-web

postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root books_api

dropdb:
	docker exec -it postgres15 dropdb books_api

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/books_api?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/books_api?sslmode=disable" -verbose down

.PHONY: build run postgres createdb dropdb migrateup migratedown