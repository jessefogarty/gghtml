package webscraper

import (
	"github.com/PuerkitoBio/goquery"
)

func Urls(doc *goquery.Document) []string {

	var urls []string

	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		link, exists := element.Attr("href")
		if exists {
			urls = append(urls, link)
			//fmt.Println(link)
		}
	})

	return urls
}