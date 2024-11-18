package db

import (
	"QueensScorecard/messages"
	"context"
	"database/sql"
	_ "embed"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

//go:embed schema.sql
var ddl string

func CreateSchema(ctx *context.Context, db *sql.DB) {
	if _, err := db.ExecContext(*ctx, ddl); err != nil {
		log.Fatal(err)
	}
}

func RunImport() {
	ctx := context.Background()
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	CreateSchema(&ctx, db)
	queries := New(db)
	for score := range messages.GetScoresFromExport() {
		err := queries.CreateScore(ctx, CreateScoreParams{
			Name:           score.Name,
			Gamenumber:     int32(score.GameNumber),
			Secondstosolve: int32(score.SecondsToSolve),
			Timestamp:      int32(score.Time.Unix()),
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(score)

	}
}
