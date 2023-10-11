DB_URL=postgresql://root:secret@localhost:5432/tudu?sslmode=disable

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

init_schema:
	migrate create -ext sql -dir db/migration -seq init_schema


createdb:
	docker exec -it postgres12 createdb --username=root --owner=root tudu

dropdb:
	docker exec -it postgres12 dropdb tudu

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server