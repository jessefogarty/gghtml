package main

import (
	"github.com/jessefogarty/goscraper/webscraper"
)

func main() {

	// TODO: Test w/ slice of urls and goroutines

	input := []string{
		"https://www.cbc.ca/news/canada/ottawa/ottawa-police-interprovincial-crossings-covid-19-1.5992615",
		"https://www.cbc.ca/news/canada/toronto/ontario-ford-long-term-care-criminal-charges-1.6016274",
	}

	// initialize source with html field
	/*source := webscraper.Webpage{
		Html: webscraper.Fetch(cbc_article),
	}*/

	for l in len(input) {}
}
