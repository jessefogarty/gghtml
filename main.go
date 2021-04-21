package main

import (
	"webscraper/fetch"
)

func main() {

	var cbc_article string = "https://www.cbc.ca/news/canada/ottawa/ottawa-police-interprovincial-crossings-covid-19-1.5992615"

	fetch.Fetch(cbc_article)

}
