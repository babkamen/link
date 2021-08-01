package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {

	type test struct {
		inputFile string
		want      map[string]string
	}

	tests := []test{
		{inputFile: "ex1.html", want: map[string]string{"/other-page": "A link to another page"}},
		{inputFile: "ex2.html", want: map[string]string{
			"https://www.twitter.com/joncalhoun": "Check me out on twitter",
			"https://github.com/gophercises":     "Gophercises is on Github!",
		}},
		{inputFile: "ex3.html", want: map[string]string{
			"https://www.twitter.com/joncalhoun": "Check me out on twitter",
		}},
	}

	for _, tc := range tests {
		file, err := os.Open(tc.inputFile)
		if err != nil {
			t.Error("Error while reading file", err)
		}
		got := Parse(file)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
