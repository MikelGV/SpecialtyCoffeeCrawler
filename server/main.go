package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/crawler"
	"github.com/gin-gonic/gin"
)
type Data struct {
    Name string
    Img string
    Price string
    Url string
}

func readData(c *gin.Context) {
    content, err := os.ReadFile("./products.json")

    if err != nil {
        log.Fatal("Error when opening file:", err)
    }

    var payload []Data
    err = json.Unmarshal(content, &payload)
    if err != nil {
        log.Fatal("Error during Unmarshal():", err)
    }
    c.IndentedJSON(http.StatusOK, payload)
}

func main() {
    r := gin.Default()
    
    timer := time.NewTicker(86400 * time.Second)
    quit := make(chan bool)
    go func() {
        for {
            select {
            case <- timer.C:
                crawler.Crawl()
            case <- quit:
                timer.Stop()
                return
            }
        }
    }()

    time.Sleep(10 * time.Second)




    r.GET("/products", readData)

    r.Run("localhost:8080")
    
}
