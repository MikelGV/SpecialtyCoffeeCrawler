package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sync"
)

type Fetcher interface {
    Fetch(url string) (body string, urls []string, err error)
}

type SafeCounter struct {
    v map[string]bool
    mu sync.Mutex
    wg sync.WaitGroup
}

var c SafeCounter = SafeCounter{v: make(map[string]bool)}

func (s SafeCounter) checkvisited(url string) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    _, ok :=s.v[url]
    if ok == false {
        s.v[url] = true
        return false
    }
    return true
}

func Crawl(url string, depth int, fetcher Fetcher) {
    defer c.wg.Done()
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
        c.wg.Add(1)
        go Crawl(u, depth-1, fetcher)
    }
    return
}

func main() {
    c.wg.Add(1)
    Crawl("https://www.reddit.com/", 4, fetcher)
    c.wg.Wait()
}

// TODO: build a fetcher that returns http results
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
    body string
    urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error){
    res, err := http.Get(url)

    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    scanner := bufio.NewScanner(res.Body)
    for i := 0; scanner.Scan() && i < 10; i++ {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }
    return "", nil, fmt.Errorf("not fount: %s", url)
}

var fetcher = fakeFetcher{}

