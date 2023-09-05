/****  Simple example with timeout  ***/

package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0} // error response
	} else {
		t := time.Since(start).Round(time.Millisecond)
		//log.Println(t)
		ch <- result{url, nil, t} // normal response
		resp.Body.Close()
	}
}

const tickRate = 2 * time.Second

func main() {
	stopper := time.After(5 * time.Second) // After funkcija koja vraca chan tj. vreca poruku kada definisano vreme istekne
	results := make(chan result)           // channel for results
	list := []string{"https://amazon.com", "https://google.com", "https://nytimes.com", "https://wsj.com"}

	for _, url := range list {
		go get(url, results) // start a CSP process
	}
	for range list { // read from the channel
		select {
		case r := <-results:
			log.Printf("%-20s %s\n", r.url, r.latency)
		case t := <-stopper: // kad primimo informaciju od stopera ispisujemo timeout
			log.Fatalf("timeout %s", t)
		}
	}
}
