package get_url

import (
	"log"
	"net/http"
	"testing"
	"time"
)

// Youtube: @mattkdvb5154
// Title: Go Class: 23 CSP, Goroutines, and Channels
// https://www.youtube.com/watch?v=zJd7Dvg3XCk&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=23

type result struct {
	url     string
	err     error
	latency time.Duration
}

// parallel
func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func TestGetUrl(t *testing.T) {
	results := make(chan result)
	list := []string{
		"https://github.com",
		"https://google.com",
		"https://youtube.com",
		"https://detik.com",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {
		r := <-results

		if r.err != nil {
			log.Printf("%-20s %s \n", r.url, r.err)
		} else {
			log.Printf("%-20s %s \n", r.url, r.latency)
		}
	}

	// equal or ...
	close(results)
}
