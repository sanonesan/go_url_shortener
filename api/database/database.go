package database

import (
	"context"
	"os"

	surrealdb "github.com/surrealdb/surrealdb.go"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *surrealdb.DB {
	surDB, _ := surrealdb.New(
		os.Getenv("DB_ADDR"),
	)
	return surDB
}
