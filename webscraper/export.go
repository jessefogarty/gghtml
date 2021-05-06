package webscraper

import (
	"encoding/json"
	"fmt"
)

func Export(articles []string) WebscraperJSON {

	for _, url := range articles {
		source := Webpage{}

		doc := Fetch(url)
		source.Links = Urls(doc)
		source.Html, _ = doc.Html()

		Metadata(doc)

		j, err := json.Marshal(source)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(j))

	}
}
