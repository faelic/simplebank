package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/faelic/simplebank/api"
	db "github.com/faelic/simplebank/db/sqlc"
	"github.com/faelic/simplebank/db/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config:", err)
	}

	if port := os.Getenv("PORT"); port != "" {
		config.ServerAddress = fmt.Sprintf("0.0.0.0:%s", port)
	}

	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal("could not connect to database:", err)
	}
	defer connPool.Close()

	store := db.NewStore(connPool)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("could not start server", err)
	}
}
