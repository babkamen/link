package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("ex1.html")
	logFatal("Error while reading file", err)

	links := Parse(file)
	fmt.Println("Links ", links)
}

func Parse(file *os.File) (links map[string]string) {
	result := make(map[string]string)
	z := html.NewTokenizer(file)
	var found bool
	var link string
	var text string
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			fmt.Println("EOF")
			return result
		case html.TextToken:
			if found {
				value := strings.Trim(string(z.Text()), "\n")
				if len(strings.TrimSpace(value)) > 0 {
					text += value
					fmt.Println("Text", value)
				}
			}
		case html.StartTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				found = true
				attr, val, _ := z.TagAttr()
				if string(attr) == "href" {
					link = string(val)
				}
			}
		case html.EndTagToken:
			tn, _ := z.TagName()
			if len(tn) == 1 && tn[0] == 'a' {
				text = strings.TrimLeft(text, "\n ")
				text = strings.TrimRight(text, "\n ")
				result[link] = text
				found = false
				link = ""
				text = ""
			}

		}
	}
}

func logFatal(message string, err error) {
	if err != nil {
		log.Fatal(message, " ", err)
	}
}
