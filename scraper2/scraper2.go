package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	args := os.Args[:]

	ch := make(chan string, 1) // input from GET
	ch2 := make(chan string, 1)

	if len(args) == 2 {

		s := strings.Split(args[1], ",")

		for i, v := range s {
			if i != len(s) {
				go get(v, ch)
				go out(ch, ch2)
				fmt.Println(<-ch2)
			}

		}

	}
	//close(ch)
}

func out(ch chan string, ch2 chan string) {
	data := <-ch
	ch2 <- data
}

func get(url string, ch chan string) {

	resp, e := http.Get(url)

	if e != nil {
		ch <- "0"
	}

	defer resp.Body.Close()

	doc, e := goquery.NewDocumentFromReader(resp.Body)

	if e != nil {
		ch <- "0"
	}

	html, _ := doc.Html()
	ch <- html

}
