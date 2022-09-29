package main

import (
	"log"
	"net/http"
	"time"
)

type Result struct {
	url      string
	err      error
	interval time.Duration
}

func GetData(url string, ch chan Result) {
	start := time.Now()
	_, err := http.Get(url)

	if err != nil {
		log.Printf("%v", err)
		result := Result{url, err, 0}
		ch <- result
	}
	t := time.Since(start).Round(time.Millisecond)
	result := Result{url, err, t}
	ch <- result

}

func main() {
	ch := make(chan Result)
	lists := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
		"https://wsaj.com",
	}
	for _, url := range lists {
		go GetData(url, ch)
	}
	for range lists {
		r := <-ch
		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		}
		log.Printf("%-20s %s", r.url, r.interval)
	}
}
