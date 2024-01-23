package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func main() {
    c := colly.NewCollector()

    c.SetRequestTimeout(120 * time.Second)
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Got a response from", r.Request.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        fmt.Println("Got this error:", err)
    })


    c.Visit("https://thecoffeestore.ie/collections/coffee")
}
