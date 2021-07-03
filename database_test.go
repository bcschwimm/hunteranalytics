package main

import "testing"

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
