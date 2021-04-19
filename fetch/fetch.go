// fetch retrieves the body of a given URL.
package fetch

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Fetches a url and returns the response body as a string
// t := Fetch("https://cbc.ca")
func Fetch(url string) string {

	// Make HTTP Request
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() // Close connection

	rb, _ := ioutil.ReadAll(resp.Body) // pass response body to std

	ctx := string(rb) // convert bytes(body) to string(ctx)

	return ctx
}
