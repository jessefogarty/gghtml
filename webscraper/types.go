package webscraper


// the main (return) object for webscraper
type Webpage struct {
	//Title       *string
	//Description *string
	InputOrder  int
	Html        string
}

type WebscraperJSON string

type ExportData struct {
	Articles []Webpage
}