package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/crawler"
	"github.com/gin-gonic/gin"
)
type Data struct {
    Id int
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

    r.GET("/crawl", func(ctx *gin.Context) {
        crawler.Crawl()
        ctx.JSON(200, gin.H{
            "msg": "this worked",
        })
    })

    r.GET("/products", readData)

    r.Run("localhost:8080")
    
}
