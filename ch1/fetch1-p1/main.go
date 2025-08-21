// Copyright © 2025 Johann Rudi

// Fetch the content found at each specified URL and write the result to a file.
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	log.SetFlags(0)              // no timestamps for a simple CLI
	log.SetPrefix("fetch1-p1: ") // prepend all log messages

	if len(os.Args) < 2 {
		log.Fatalf("usage: %s url ...", os.Args[0])
	}

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("GET %s: %v", url, err)
		}
		defer resp.Body.Close()

		filename := sanitizeFilename(url)
		out, err := os.Create(filename)
		if err != nil {
			log.Fatalf("creating file %s: %v", filename, err)
		}
		defer out.Close()

		if _, err := io.Copy(out, resp.Body); err != nil {
			log.Fatalf("writing to %s: %v", filename, err)
		}

		log.Printf("saved %s to %s", url, filename)
	}
}

func sanitizeFilename(url string) string {
	f := strings.ReplaceAll(url, "/", "_")
	f = strings.ReplaceAll(f, ":", "_")
	return f
}
