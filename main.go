package main

import (
	"encoding/json"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/jessefogarty/goscraper/webscraper"
)

func main() {

	var cbc_article string = "https://www.cbc.ca/news/canada/ottawa/ottawa-police-interprovincial-crossings-covid-19-1.5992615"

	// initialize source with html field
	/*source := webscraper.Webpage{
		Html: webscraper.Fetch(cbc_article),
	}*/

	var source struct {
		Title       string
		Description string
		Html        string
		Links       []string
		doc         *goquery.Document
	}

	source.doc = webscraper.Fetch(cbc_article)
	source.Links = webscraper.Urls(source.doc)
	source.Html, _ = source.doc.Html()

	webscraper.Metadata(source.doc)

	j, _ := json.Marshal(source)

	fmt.Printf("%s", string(j))
}
