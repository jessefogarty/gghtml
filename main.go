package main

import (
	"os"
	"strings"

	"github.com/jessefogarty/gghtml/gghtml"
)

func main() {

	args := os.Args[:]
	if len(args) == 2 {
		strin := strings.Split(args[1], ",")
		gghtml.Fetch(strin)
	} else {
		os.Exit(1)
	}

}
