package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jessefogarty/gghtml/gghtml"
)

// main() is the command line entrace to gghtml. Requires 1 string argument.
//
// splits string argument on commas and passing slice to gghtml.Fetch()
//
func main() {

	//args := os.Args[:]
	strin := strings.Split(os.Args[1], ",")
	t0 := time.Now()
	gghtml.Fetch(strin)
	t1 := time.Now()
	fmt.Printf("Finished in: %vms", t1.Sub(t0).Seconds())

}
