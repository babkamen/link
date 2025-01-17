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
			"#":                                "Login",
			"/lost":                            "Lost? Need help?",
			"https://twitter.com/marcusolsson": "@marcusolsson",
		}},
		{inputFile: "ex4.html", want: map[string]string{
			"/dog-cat": "dog cat",
		}},
	}

	for _, tc := range tests {
		file, err := os.Open(tc.inputFile)
		if err != nil {
			t.Error("Error while reading file", err)
		}
		got := convertToMap(Parse(file))
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func convertToMap(links []Link) map[string]string {
	result := make(map[string]string, len(links))
	for _, link := range links {
		result[link.Href] = link.Text
	}
	return result
}
