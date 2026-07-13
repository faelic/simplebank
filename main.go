package main

import (
	"context"
	"log"

	"github.com/faelic/simplebank/api"
	db "github.com/faelic/simplebank/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("could not connect to database:", err)
	}

	defer connPool.Close()

	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("could not start server", err)
	}
}
