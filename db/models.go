// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

type Score struct {
	ID             int64
	Name           string
	Gamenumber     int64
	Secondstosolve int64
	Timestamp      int64
}