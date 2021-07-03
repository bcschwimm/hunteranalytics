package main

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestList(t *testing.T) {
	want := []Metrics{{7, 7, 7, 7, "06/29/2021"}, {10, 10, 10, 10, "05/29/2021"}}
	got := []Metrics{}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error %s was not expected opening stub db connection", err)
	}
	defer db.Close()

	// mocking our sql rows to test that list is parsing returned rows correctly into []Metrics{}
	rows := sqlmock.NewRows([]string{"playing", "training", "exercising", "woofing", "date"}).
		AddRow(7, 7, 7, 7, "06/29/2021").
		AddRow(10, 10, 10, 10, "05/29/2021")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)
	rs, _ := db.Query("SELECT")

	for rs.Next() {
		var testMetric Metrics
		rs.Scan(&testMetric.Playing, &testMetric.Training, &testMetric.Exercising, &testMetric.Woofing, &testMetric.Date)
		got = append(got, testMetric)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Unexpected Metrics Struct output for List got %v want %v", got, want)
	}
}

func TestIntConv(t *testing.T) {
	var tests = []struct {
		input string
		want  int
	}{
		{"42", 42},
		{"-42", -42},
		{"200", 200},
		{"0", 0},
		{"1", 1},
		{"9", 9},
		{"9.9", 0},
		{"pizza", 0},
	}
	for _, test := range tests {
		if got := intConv(test.input); got != test.want {
			t.Errorf("Int Conversion Test Failed got %v want %v", got, test.want)
		}
	}
}
