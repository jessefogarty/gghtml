package main

import (
	"encoding/json"
	"fmt"

	"github.com/jessefogarty/goscraper/webscraper"
)

func main() {

	// TODO: Test w/ slice of urls and goroutines

	var cbc_article string = "https://www.cbc.ca/news/canada/ottawa/ottawa-police-interprovincial-crossings-covid-19-1.5992615"

	// initialize source with html field
	/*source := webscraper.Webpage{
		Html: webscraper.Fetch(cbc_article),
	}*/


	source := webscraper.Webpage{}

	doc := webscraper.Fetch(cbc_article)
	source.Links = webscraper.Urls(doc)
	source.Html, _ = doc.Html()

	webscraper.Metadata(doc)

	j, err := json.Marshal(source)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
}
