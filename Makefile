postgres:
	docker run --name postgres18 -p 5432:5432 -e POSTGRES_USER=favour -e POSTGRES_PASSWORD=faelicdika -d postgres:18-bookworm
createdb:
	docker exec -it postgres18 createdb --username=favour --owner=favour simple_bank

dropdb:
	docker exec -it postgres18 dropdb --username=favour --if-exists simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/faelic/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock