package webscraper

import (
	"github.com/PuerkitoBio/goquery"
)

// the main (return) object for webscraper
type Webpage struct {
	Title       *string
	Description *string
	Html        *goquery.Document
	Links       []string
}
