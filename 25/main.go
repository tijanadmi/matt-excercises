/****  Simple example with timeout  ***/

package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(ctx context.Context, url string, ch chan<- result) {
	start := time.Now()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if resp, err := http.DefaultClient.Do(req); err != nil { // sada smo definisali kontekst 5 sec u koliko nema response prekida se poziv i prosledjuje se error
		ch <- result{url, err, 0} // error response -- kada kontekst istekne dobijamo ovde error: context deadline exceeded
	} else {
		t := time.Since(start).Round(time.Millisecond)
		//log.Println(t)
		ch <- result{url, nil, t} // normal response
		resp.Body.Close()
	}
}

func main() {

	results := make(chan result) // channel for results
	list := []string{"https://amazon.com", "https://google.com", "https://nytimes.com", "https://wsj.com"}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	for _, url := range list {
		go get(ctx, url, results) // start a CSP process
	}
	for range list { // read from the channel
		r := <-results
		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)
		}

	}
}
