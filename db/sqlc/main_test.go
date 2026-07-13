package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/faelic/simplebank/db/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testStore Store
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	ctx := context.Background()

	connPool, err := pgxpool.New(ctx, config.DBSource)
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
