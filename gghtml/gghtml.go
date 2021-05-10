package gghtml

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// Handles the threaded downloading of HTML source for each URL in the array argument.
//
// Returns string JSON object (of ExportData{}) to stdout
func Fetch(s []string) {

	t0 := time.Now()

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

	t1 := time.Now()

	fmt.Printf("%s \n %v", j, t1.Sub(t0))
}

/*
Handles the HTTP GET request for a given URL string.

Creates a Webpage{} object with retained input order and html.

Adds the new Webpage{} to the buffered channel.
*/
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
