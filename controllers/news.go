package controllers

import (
	"GO-SCRAPING/structs"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (idb *InDB) ScrapeAllNews(c *gin.Context) {
	var (
		result gin.H
	)

	res, err := http.Get("https://covid19.go.id/p/berita")
	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	rows := make([]structs.News, 0)

	doc.Find(".row .col-lg-4").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(structs.News)
		row.Title = sel.Find("h4").Text()
		row.Image, _ = sel.Find("img").Attr("data-original")
		row.Link, _ = sel.Find("a").Attr("href")
		row.PostDate = sel.Find("time").Text()
		rows = append(rows, *row)
	})

	if len(rows) == 0 {
		result = gin.H{
			"status":  404,
			"message": "News is null",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	} else {
		result = gin.H{
			"status":     200,
			"message":    "Success Retrieving News Data",
			"data_count": len(rows),
			"data":       rows,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) ScrapeDetailNews(c *gin.Context) {
	var (
		result gin.H
	)

	urlBind := c.Request.URL.Query().Get("url")

	res, err := http.Get(urlBind)
	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		result = gin.H{
			"status":  400,
			"message": "Bad Request",
		}
		c.JSON(http.StatusBadRequest, result)
		return
	}

	rows := make([]structs.NewsDetail, 0)

	doc.Find(".blog-posts").Children().Each(func(i int, sel *goquery.Selection) {
		row := new(structs.NewsDetail)
		row.Title = sel.Find("h2").Text()
		row.PostDate = sel.Find("time").Text()
		row.Description = sel.Find("p").Text()
		row.Image, _ = sel.Find("img").Attr("src")
		rows = append(rows, *row)
	})

	if len(rows) == 0 {
		result = gin.H{
			"status":  404,
			"message": "News is null",
		}

		c.JSON(http.StatusNotFound, result)
		return
	} else {
		result = gin.H{
			"status":     200,
			"message":    "Success Retrieving News Data",
			"data_count": len(rows),
			"data":       rows,
		}
	}

	c.JSON(http.StatusOK, result)
}
