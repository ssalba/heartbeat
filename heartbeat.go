package main

import (
	"flag"
	"fmt"
	"os"
)

var Version = "No version provided"

func main() {
	ver := flag.Bool("v", false, "Print version")
	flag.Parse()

	if *ver {
		fmt.Printf("%s %s\n", os.Args[0], Version)
		os.Exit(0)
	}
}
