package main

import (
	"fmt"

	"github.com/jessefogarty/goscraper/webscraper"
)

func main() {

	var cbc_article string = "https://www.cbc.ca/news/canada/ottawa/ottawa-police-interprovincial-crossings-covid-19-1.5992615"

	doc := webscraper.Fetch(cbc_article)

	a := webscraper.Urls(doc)

	fmt.Println(len(a))
}
