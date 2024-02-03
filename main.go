package main

import (
	"fmt"

	"github.com/MikelGV/SpecialtyCoffeeCrawler/crawler"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.GET("/crawl", func(ctx *gin.Context) {
        crawler.Crawl()
        fmt.Println("hello world")
        ctx.JSON(200, gin.H{
            "msg": "this worked",
        })
    })
    r.Run("localhost:8080")
    
}
