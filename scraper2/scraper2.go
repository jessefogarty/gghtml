package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	args := os.Args[:]

	if len(args) == 2 {

		s := strings.Split(args[1], ",")

		ch := make(chan Webpage, len(s)) // init buffer
		var wg sync.WaitGroup

		for i, v := range s {
			if i != len(s) {
				wg.Add(1)
				go get(v, i, ch, &wg)
			}

		}
		wg.Wait(); close(ch)

		articles := ExportData{}

		for a := range ch {

			articles.Articles = append(articles.Articles, a)

		}

		j, _ := json.Marshal(articles)

		fmt.Printf("%s", j)


	}
}

func get(url string, i int, ch chan Webpage, wg *sync.WaitGroup) {

	defer wg.Done()

	article := Webpage{InputOrder: i}

	resp, e := http.Get(url)

	if e != nil {
		article.Html = "0"
		ch <- article
	}

	defer resp.Body.Close()

	doc, e := goquery.NewDocumentFromReader(resp.Body)

	if e != nil {
		article.Html = "0"
		ch <- article
	}

	html, _ := doc.Html()
	article.Html = html

	ch <- article

}

// Webpage represents the article object for each url.
// Html source is scraped and input queue is saved.
type Webpage struct {
	InputOrder int
	Html       string
}

// Ths structure of data in transit.
// From Webpage to WebscraperJSON.
type ExportData struct {
	Articles []Webpage
}
