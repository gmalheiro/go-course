DB_URL=postgresql://root:root@localhost:5432/simple-bank?sslmode=disable

postgres:
	docker run --name db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres dropdb simple-bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc 
