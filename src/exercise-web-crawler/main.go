package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type WorkerCounter struct {
	mu    sync.Mutex
	count int
}

func (w *WorkerCounter) Inc() {
	w.mu.Lock()
	w.count++
	w.mu.Unlock()
}

func (w *WorkerCounter) Dec() {
	w.mu.Lock()
	w.count--
	w.mu.Unlock()
}

func (w *WorkerCounter) IsZero() bool {
	return w.count == 0
}

type CrawlList struct {
	mu sync.Mutex
	v  map[string]string
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	counter := WorkerCounter{}
	crawl_list := CrawlList{v: make(map[string]string)}
	ch := make(chan string)
	counter.Inc()
	go crawl(url, depth, fetcher, &counter, &crawl_list, ch)
	for v := range ch {
		fmt.Println(v)
	}
}

func crawl(url string, depth int, fetcher Fetcher, counter *WorkerCounter, crawl_list *CrawlList, ch chan string) {
	defer func() {
		counter.Dec()
		if counter.IsZero() {
			close(ch)
		}
	}()

	if depth <= 0 {
		return
	}
	if _, ok := crawl_list.v[url]; ok {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		crawl_list.mu.Lock()
		crawl_list.v[url] = ""
		crawl_list.mu.Unlock()
		fmt.Println(err)
		return
	}
	crawl_list.mu.Lock()
	crawl_list.v[url] = body
	crawl_list.mu.Unlock()
	ch <- fmt.Sprintf("found: %s %q\n", url, body)

	for _, u := range urls {
		counter.Inc()
		go crawl(u, depth-1, fetcher, counter, crawl_list, ch)
	}

	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Second)
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
