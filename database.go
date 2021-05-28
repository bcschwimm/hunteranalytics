package main

import (
	"net/url"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host       string
	Name       string
	User       string
	Password   string
	DisableTLS bool
}

// Open knows how to open a database connection
func Open(cfg Config) (*sqlx.DB, error) {
	q := url.Values{}
	q.Set("sslmode", "disable")
	if cfg.DisableTLS {
		q.Set("sslmode", "disable")
	}

	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(cfg.User, cfg.Password),
		Host:     cfg.Host,
		Path:     cfg.Name,
		RawQuery: q.Encode(),
	}
	// passing url string with scheme, user, pw etc to .Open
	return sqlx.Open("postgres", u.String())
}

// List reads all data from hunter table and returns it
func List(db *sqlx.DB) ([]Metrics, error) {
	list := []Metrics{}

	const q = `SELECT playing, training, excercising, woofing, date FROM hunter`
	if err := db.Select(&list, q); err != nil {
		return nil, err
	}
	return list, nil
}
