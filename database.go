package main

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Metrics are datapoints tracked in miniutes for hunters behavoir
type Metrics struct {
	Playing    int    `json:"playing"`
	Training   int    `json:"training"`
	Exercising int    `json:"exercising"`
	Woofing    int    `json:"woofing"`
	Date       string `json:"date"`
}

func pass() string {
	text, err := ioutil.ReadFile("conn.txt")
	if err != nil {
		panic(err.Error())
	}
	return string(text)
}

// Open opens our cloud sql connection
func Open() (*sql.DB, error) {
	dbURI := pass()
	return sql.Open("pgx", dbURI)
}

// List reads all data from hunter table and returns it
func List(db *sql.DB) ([]Metrics, error) {
	list := []Metrics{}

	const q = `SELECT playing, training, exercising, woofing, date FROM hunter`

	result, err := db.Query(q)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var metric Metrics
		err := result.Scan(&metric.Playing, &metric.Training, &metric.Exercising, &metric.Woofing, &metric.Date)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, metric)
	}

	return list, nil
}
