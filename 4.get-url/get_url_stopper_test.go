package get_url

import (
	"log"
	"net/http"
	"testing"
	"time"
)

// Youtube: @mattkdvb5154
// Title: Go Class: 24 Select
// https://www.youtube.com/watch?v=tG7gII0Ax0Q&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=24

type resultStopper struct {
	url     string
	err     error
	latency time.Duration
}

// parallel
func getStopper(url string, ch chan<- resultStopper) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- resultStopper{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- resultStopper{url, nil, t}
		resp.Body.Close()
	}
}

func TestGetUrlStopper(t *testing.T) {
	stopper := time.After(3 * time.Second)
	results := make(chan resultStopper)

	list := []string{
		"https://github.com",
		"https://google.com",
		"https://youtube.com",
		"https://detik.com",
		"http://localhost:8080",
	}

	for _, url := range list {
		go getStopper(url, results)
	}

	for range list {
		select {
		case r := <-results:
			if r.err != nil {
				log.Printf("%-20s %s \n", r.url, r.err)
			} else {
				log.Printf("%-20s %s \n", r.url, r.latency)
			}
		case t := <-stopper:
			log.Fatalf("timeout %s", t)
		}
	}

	// just the same but ...
	close(results)
}
