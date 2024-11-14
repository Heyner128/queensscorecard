package main

import (
	"QueensScorecard/cmd"
	"net/http"
)

func runApi() error {
	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		return err
	}

	return nil
}

func main() {
	cmd.Execute()
}
