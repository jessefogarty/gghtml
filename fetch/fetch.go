// Package returns the source code of
// a URL as a string.
package fetch

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Fetch GET's the URL and returns a
// goquery string document.
//
//	example:
//
// 		fetch.Fetch("https://cbc.ca")
func Fetch(url string) {

	// Make HTTP Request
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // Close connection

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Find <a> href text and print
	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		link, exists := element.Attr("href")
		if exists {
			fmt.Println(link)
		}
	})

}
