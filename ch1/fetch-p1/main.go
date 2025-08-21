// Copyright © 2025 Johann Rudi

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(0)             // no timestamps for a simple CLI
	log.SetPrefix("fetch-p1: ") // prepend all log messages

	if len(os.Args) < 2 {
		log.Fatalf("usage: %s url ...", os.Args[0])
	}

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("GET %s: %v", url, err)
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("reading %s: %v", url, err)
		}

		fmt.Printf("%s", b)
	}
}
