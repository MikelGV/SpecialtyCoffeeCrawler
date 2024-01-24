package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)
type Product struct {
   Name string
   Img string
   Price string
   Url string
}

func main() {
    c := colly.NewCollector()
    c.SetRequestTimeout(120 * time.Second)
    products := make([]Product, 0)

    // TODO: make it scrape different divs
    c.OnHTML("div.product-loop", func(h *colly.HTMLElement) {
        h.ForEach("div.product-index", func(i int, h *colly.HTMLElement) {
            item := Product{}
            item.Name = h.Text
            item.Img = h.ChildAttr("div.ci", "a")
            item.Price = h.Attr("prod-price")
            item.Url = "https://thecoffeestore.ie/collections/coffee/products" + h.Attr("href")
            products = append(products, item)
        })
    })


    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL)
    })

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Got a response from", r.Request.URL)
    })

    c.OnError(func(r *colly.Response, err error) {
        fmt.Println("Got this error:", err)
    })


    c.OnScraped(func(r *colly.Response) {
        fmt.Println("Finished", r.Request.URL)
        js, err := json.MarshalIndent(products, "", " ")
        if err != nil {
            log.Fatal(err)
        }
        
        if err := os.WriteFile("products.json", js, 0664); err == nil {
            fmt.Println("Data written to file successfully")
        }

    })

    c.Visit("https://thecoffeestore.ie/collections/coffee")
}
