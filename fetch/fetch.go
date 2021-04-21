// fetch retrieves the body of a given URL.
package fetch

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// func extractLink(index int, element *goquery.Selection) {}

// Fetches a url and returns the response body as a string
// t := Fetch("https://cbc.ca")
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

	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		link, exists := element.Attr("href")
		if exists {
			fmt.Println(link)
		}
	})
}
