package api

import (
	"QueensScorecard/db"
	"context"
	"database/sql"
	_ "embed"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type FastestByMonthDto struct {
	Nombre string
	Meses  []string
}

type FastestByWeekDto struct {
	Nombre  string
	Semanas []string
}

type ScoreDto struct {
	Nombre               string
	NumeroDeJuego        int
	SegundosParaResolver int
	Fecha                time.Time
}

func scores(c *gin.Context) {
	ctx := context.Background()
	database, err := sql.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	queries := db.New(database)
	scores, err := queries.GetScores(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	var response []ScoreDto
	for _, v := range scores {
		response = append(response, ScoreDto{
			Nombre:               v.Name,
			NumeroDeJuego:        int(v.Gamenumber),
			SegundosParaResolver: int(v.Secondstosolve),
			Fecha:                time.Unix(int64(v.Timestamp), 0),
		})
	}
	c.JSON(http.StatusOK, response)
}

func topByWeek(c *gin.Context) {
	ctx := context.Background()
	database, err := sql.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	queries := db.New(database)
	fastestByMonth, err := queries.GetFastestPlayersByWeek(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	var response []FastestByWeekDto
	for _, v := range fastestByMonth {
		response = append(response, FastestByWeekDto{
			Nombre:  v.Name,
			Semanas: strings.Split(v.Week.String, ","),
		})
	}
	c.JSON(http.StatusOK, response)
}

func topByMonth(c *gin.Context) {
	ctx := context.Background()
	database, err := sql.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	queries := db.New(database)
	fastestByMonth, err := queries.GetFastestPlayersByMonth(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	var response []FastestByMonthDto
	for _, v := range fastestByMonth {
		response = append(response, FastestByMonthDto{
			Nombre: v.Name,
			Meses:  strings.Split(v.Months.String, ","),
		})
	}
	c.JSON(http.StatusOK, response)
}

func RunApi() {
	r := gin.Default()
	r.GET("/mejoresPorMes", topByMonth)
	r.GET("/mejoresPorSemana", topByWeek)
	r.GET("/puntajes", scores)
	ctx := context.Background()
	database, err := sql.Open("mysql", os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	db.CreateSchema(&ctx, database)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
