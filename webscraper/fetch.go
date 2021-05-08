package webscraper

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Fetch GET's the URL and returns a
// goquery string document.
//
//	example:
//
// 	fetch.Fetch("https://cbc.ca")
func Fetch(url *string) string{

	// Make HTTP Request
	resp, err := http.Get(*url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // Close connection

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	html, _ := doc.Html()

	return html

}

