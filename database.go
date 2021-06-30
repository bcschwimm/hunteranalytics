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

const (
	sqlinsert = `INSERT INTO hunter (playing, training, exercising, woofing, date)
VALUES ($1, $2, $3, $4, $5)`
	q = `SELECT playing, training, exercising, woofing, date FROM hunter`
)

func pass() string {
	text, err := ioutil.ReadFile("conn.txt")
	if err != nil {
		panic(err.Error())
	}
	return string(text)
}

// insert submits metric struct data into our database
func (m Metrics) insert() error {
	db, err := Open()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec(sqlinsert, m.Playing, m.Training, m.Exercising, m.Woofing, m.Date)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

// Open opens our cloud sql connection
func Open() (*sql.DB, error) {
	dbURI := pass()
	return sql.Open("pgx", dbURI)
}

// List reads all data from hunter table and returns it
func List(db *sql.DB) ([]Metrics, error) {
	list := []Metrics{}

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
