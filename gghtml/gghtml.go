package gghtml

import (
	//"encoding/json"
	//"fmt"
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"
)

// Handles the threaded downloading of HTML source for each URL in the array argument.
//
// Returns string JSON object (of ExportData{}) to stdout
func Fetch(s []string) (*[]byte, float64) {

	t0 := time.Now()
	ch := make(chan Webpage, len(s)) // init buffer
	
	var wg sync.WaitGroup
	articles := &ExportData{}

	for i, v := range s {
		//wg.Add(1)
		go get(v, i, &ch, &wg)
		
	}
	wg.Wait()
	close(ch)

	
	for a := range ch {
		articles.Articles = append(articles.Articles, &a)
		wg.Done()
	}
	
	j, _ := json.Marshal(articles)
	t1 := time.Now()
	return &j, t1.Sub(t0).Seconds()

}

/*
Handles the HTTP GET request for a given URL string.

Creates a Webpage{} object with retained input order and html.

Adds the new Webpage{} to the buffered channel.
*/
func get(url string, i int, ch *chan Webpage, wg *sync.WaitGroup) {

	wg.Add(1)

	article := &Webpage{InputOrder: i}

	resp, _ := http.Get(url)

	html, _ := io.ReadAll(resp.Body)

	resp.Body.Close()

	h := string(html)

	article.Html = h

	*ch <- *article

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
	Articles []*Webpage
}
