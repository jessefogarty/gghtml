package webscraper


// Webpage represents the article object for each url.
// Html source is scraped and input queue is saved.
type Webpage struct {
	InputOrder  int
	Html        string
}

// A string representation of all processed data.
type WebscraperJSON string

// Ths structure of data in transit.
// From Webpage to WebscraperJSON.
type ExportData struct {
	Articles []Webpage
}