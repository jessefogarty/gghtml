package webscraper

import (
	"fmt"
	"regexp"
	"strings"

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

func Title(doc *goquery.Document) {
	doc.Find("meta").Each(func(index int, element *goquery.Selection) {
		tag, _ := element.Attr("property")
		if strings.Contains(tag, "og"){
			fmt.Println(tag)
		}
	})
}
