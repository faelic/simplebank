postgres:
	docker run --name monierave-postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=favour -e POSTGRES_PASSWORD=faelicdika -d postgres:18-bookworm
createdb:
	docker exec -it monierave-postgres createdb --username=favour --owner=favour simple_bank

dropdb:
	docker exec -it monierave-postgres dropdb --username=favour --if-exists simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

migrateup2:
	migrate -path db/migration -database "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable" -verbose up 2

migratedown2:
	migrate -path db/migration -database "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable" -verbose down 2

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/faelic/monierave/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc server mock
