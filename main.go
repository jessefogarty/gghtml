package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessefogarty/goscraper/webscraper"
)

func main() {

	args := os.Args[:]
	if len(args) == 2 {
		str := strings.Split(args[1], ",")
		fmt.Println(str)
		fmt.Printf("%s", webscraper.Scraper(str))
	} else {
		os.Exit(1)
	}

}
