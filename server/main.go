package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/crawler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
    
    err := godotenv.Load()

    if err != nil {
        log.Fatal("Error loading .env file:", err)
    }

    port := os.Getenv("PORT")
    
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

    r.Run(port)
    
}
