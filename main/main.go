package main

import (
	"GO-SCRAPING/config"
	"GO-SCRAPING/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()
	router.Use(cors.Default())

	news := router.Group("/news")
	{
		news.GET("/", inDB.ScrapeAllNews)
		news.GET("/limit", inDB.ScrapeLimitNews)
		news.GET("/detail", inDB.ScrapeDetailNews)
	}

	err := router.Run(":47000")

	if err != nil {
		panic("Error when running router")
	}
}
