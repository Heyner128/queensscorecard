package db

import (
	"QueensScorecard/messages"
	"context"
	"database/sql"
	_ "embed"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

func RunImport() error {
	ctx := context.Background()
	db, err := sql.Open("sqlite3", "file:scores.db?mode=rwc")
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
			Gamenumber:     int64(score.GameNumber),
			Secondstosolve: int64(score.SecondsToSolve),
			Timestamp:      score.Time.Unix(),
		})
		if err != nil {
			return err
		}

	}
	return nil
}
