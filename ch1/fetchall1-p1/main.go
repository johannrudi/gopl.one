// Copyright © 2025 Johann Rudi

// Fetchall fetches URLs in parallel and and writes the results to files.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	log.SetFlags(0)                // no timestamps for a simple CLI
	log.SetPrefix("fetchall-p1: ") // prepend all log messages

	if len(os.Args) < 2 {
		log.Fatalf("usage: %s url ...", os.Args[0])
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
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

	nbytes, err := io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("writing to %s: %v", filename, err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  saved %s to %s", secs, nbytes, url, filename)
}

func sanitizeFilename(url string) string {
	f := strings.ReplaceAll(url, "/", "_")
	f = strings.ReplaceAll(f, ":", "_")
	return f
}
