package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jessefogarty/gghtml/gghtml"
)

// main() is the command line entrace to gghtml. Requires 1 string argument.
//
// splits string argument on commas and passing slice to gghtml.Fetch()
//
func main() {

	//args := os.Args[:]
	strin := strings.Split(os.Args[1], ",")
	_, t := gghtml.Fetch(strin)
	fmt.Printf("Golang finished in: \n %f", t)

}