package main

import (
	"os"
	"strings"

	"github.com/jessefogarty/gghtml/gghtml"
)


// main() is the command line entrace to gghtml. Requires 1 string argument.
//
// splits string argument on commas and passing slice to gghtml.Fetch()
//	
func main() {

	args := os.Args[:]
	if len(args) == 2 {
		strin := strings.Split(args[1], ",")
		gghtml.Fetch(strin)
	} else {
		os.Exit(1)
	}

}
