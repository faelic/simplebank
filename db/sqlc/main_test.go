package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testStore *Store
var testDB *pgxpool.Pool

const dbSource = "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable"

func TestMain(m *testing.M) {
	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, dbSource)
	if err != nil {
		log.Fatal("could not connect to database:", err)
	}

	defer connPool.Close()

	testDB = connPool
	testQueries = New(testDB)
	testStore = NewStore(testDB)

	code := m.Run()
	testDB.Close()

	os.Exit(code)
}
