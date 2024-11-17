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

func RunImport() error {
	ctx := context.Background()
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		return err
	}
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		return err
	}
	queries := New(db)
	for score := range messages.GetScoresFromExport() {
		err := queries.CreateScore(ctx, CreateScoreParams{
			Name:           score.Name,
			Gamenumber:     int32(score.GameNumber),
			Secondstosolve: int32(score.SecondsToSolve),
			Timestamp:      int32(score.Time.Unix()),
		})
		if err != nil {
			return err
		}
		log.Println(score)

	}
	return nil
}
