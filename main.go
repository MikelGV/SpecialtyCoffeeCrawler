package main

import (
	"fmt"
	"sync"
    "net/http"
)

type Fetcher interface {
    Fetch(url string) (body string, urls []string, err error)
}

type SafeCounter struct {
    v map[string]bool
    mu sync.Mutex
}

var c SafeCounter = SafeCounter{v: make(map[string]bool)}

func Crawl(url string, depth int, fetcher Fetcher) {
    if depth <=0 {
        return
    }

    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        Crawl(u, depth-1, fetcher)
    }
    return
}

func main() {

    Crawl("https://www.reddit.com/", 4, fetcher)
}

// TODO: build a fetcher that returns http results

func Fetch(url string) (string, []string, error) {
    res, err := http.Get(url)
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()
}


