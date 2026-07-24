package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/faelic/monierave/db/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testStore Store
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Printf("skipping db/sqlc tests: cannot load config: %v", err)
		os.Exit(0)
	}
	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Printf("skipping db/sqlc tests: could not create database pool: %v", err)
		os.Exit(0)
	}

	if err := connPool.Ping(ctx); err != nil {
		log.Printf("skipping db/sqlc tests: database is unreachable: %v", err)
		os.Exit(0)
	}

	defer connPool.Close()

	testDB = connPool
	testQueries = New(testDB)
	testStore = NewStore(testDB)

	code := m.Run()
	testDB.Close()

	os.Exit(code)
}
