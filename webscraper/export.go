package webscraper

import (
	"encoding/json"
	"fmt"
)

func Export(articles ExportData) string {

	j, err := json.Marshal(articles)

	if err != nil {
		fmt.Println(err)
	}

	return string(j)

}
