package webscraper

import (
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

func Urls(doc *goquery.Document) []string {

	var urls []string

	re := regexp.MustCompile(`^\w`) // could get <domain>.com onwards. Finall.

	doc.Find("a").Each(func(index int, element *goquery.Selection) {
		link, exists := element.Attr("href")
		if exists && re.MatchString(link) {
			urls = append(urls, link)
			//fmt.Println(link)
		}
	})

	return urls
}

func Title(doc *goquery.Document) string {
	
}