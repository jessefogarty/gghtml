package webscraper

func convert(ch chan Webpage, l []string) {

	for i := range l {
		ch <- Webpage{Html: Fetch(l[i]), InputOrder: i}
	}

	close(ch)
}

func Scraper(input []string) WebscraperJSON {
	ch := make(chan Webpage)

	articles := ExportData{}

	go convert(ch, input)

	for a := range ch {
		//time.Second.Milliseconds(30)
		articles.Articles = append(articles.Articles, a)
	}

	data := WebscraperJSON(Export(articles))


	return data

}
