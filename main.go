package main

import (
	"fmt"

	"github.com/jessefogarty/goscraper/webscraper"
)

func main() {

	var cbc_article string = "https://www.cbc.ca/news/canada/ottawa/ottawa-police-interprovincial-crossings-covid-19-1.5992615"

	// initialize source with html field
	source := webscraper.Webpage{
		Html: webscraper.Fetch(cbc_article),
	}

	source.Links = webscraper.Urls(source.Html)

	fmt.Println(source)
}
