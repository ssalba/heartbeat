package main

import (
	"flag"
	"fmt"
	"net/http"
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

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":5000", nil)
}
