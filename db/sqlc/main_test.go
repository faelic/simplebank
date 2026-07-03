package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

var testQueries *Queries

const dbSource = "postgresql://favour:faelicdika@localhost:5432/simple_bank?sslmode=disable"

func TestMain(m *testing.M) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dbSource)
	if err != nil {
		log.Fatal("could not connect to database:", err)
	}

	testQueries = New(conn)

	code := m.Run()
	conn.Close(ctx)

	os.Exit(code)
}
