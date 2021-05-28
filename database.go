package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// Open opens our cloud sql connection
func Open() (*sql.DB, error) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", "35.245.113.254", "postgres", "hunter_db_420!", "5432", "hunterdb")
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
