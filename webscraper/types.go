package webscraper

import (
	"github.com/PuerkitoBio/goquery"
)

// the main (return) object for webscraper
type Source struct {
	title       *string
	description *string
	html        *goquery.Document
	links       []*string
}
