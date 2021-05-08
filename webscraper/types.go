package webscraper


// the main (return) object for webscraper
type Webpage struct {
	InputOrder  int
	Html        string
}

type WebscraperJSON string

type ExportData struct {
	Articles []Webpage
}