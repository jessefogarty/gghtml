package gghtml

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func Fetch(s []string) {

	ch := make(chan Webpage, len(s)) // init buffer
	var wg sync.WaitGroup

	for i, v := range s {
		if i != len(s) {
			wg.Add(1)
			go get(v, i, ch, &wg)
		}

	}
	wg.Wait()
	close(ch)

	articles := ExportData{}

	for a := range ch {

		articles.Articles = append(articles.Articles, a)

	}

	j, _ := json.Marshal(articles)

	fmt.Printf("%s", j)

}

// get performs an HTTP GET on the url string
// the
func get(url string, i int, ch chan Webpage, wg *sync.WaitGroup) {

	defer wg.Done()

	article := Webpage{InputOrder: i}

	resp, e := http.Get(url)

	if e != nil {
		article.Html = "0"
		ch <- article
	}

	html, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	//doc, e := goquery.NewDocumentFromReader(resp.Body)

	if e != nil {
		article.Html = "0"
		ch <- article
	}

	//html, _ := doc.Html()
	article.Html = string(html)

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
