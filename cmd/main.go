package main

import (
	// "log"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Crafter-deo/tech-trends-api/src/websites"
	"github.com/gin-gonic/gin"
)

func main() {
	// log.Println(websites.ScrapeCnet())
	router := gin.Default()

	router.GET("/all", func(ctx *gin.Context) {
		all_blogs := [][]websites.Blogs{}
		for _, site := range websites.Sites {
			switch site {
			case "cnet":
				blogs := websites.ScrapeCnet()
				all_blogs = append(all_blogs, blogs)

			case "codingdojo":
				blogs := websites.ScrapeCodingdojo()
				all_blogs = append(all_blogs, blogs)

			case "digitaltrends":
				blogs := websites.ScrapeDigitaltrends()
				all_blogs = append(all_blogs, blogs)

			case "hackernews":
				blogs := websites.ScrapeHackernews()
				all_blogs = append(all_blogs, blogs)

			case "mashable":
				blogs := websites.ScrapeMashable()
				all_blogs = append(all_blogs, blogs)

			case "medium":
				blogs := websites.ScrapeMedium()
				all_blogs = append(all_blogs, blogs)
			}
		}
		ctx.JSON(http.StatusOK, all_blogs)
	})

	router.GET("/cnet", func(ctx *gin.Context) {
		blogs := websites.ScrapeCnet()
		ctx.JSON(http.StatusOK, blogs)
	})
	// TODO: not returning blogs, returns null, check scraper
	router.GET("/codingdojo", func(ctx *gin.Context) {
		blogs := websites.ScrapeCodingdojo()
		ctx.JSON(http.StatusOK, blogs)
	})
	// TODO: not returning blogs, returns nothing, check scraper
	router.GET("/digitaltrends", func(ctx *gin.Context) {
		blogs := websites.ScrapeDigitaltrends()
		ctx.JSON(http.StatusOK, blogs)
	})
	router.GET("/hackernews", func(ctx *gin.Context) {
		blogs := websites.ScrapeHackernews()
		ctx.JSON(http.StatusOK, blogs)
	})
	router.GET("/mashable", func(ctx *gin.Context) {
		blogs := websites.ScrapeMashable()
		ctx.JSON(http.StatusOK, blogs)
	})
	// TODO: not returning blogs, returns null, check scraper
	router.GET("/medium", func(ctx *gin.Context) {
		blogs := websites.ScrapeMedium()
		ctx.JSON(http.StatusOK, blogs)
	})
	router.Run()
}


func loadSites() []string {
	file, err := os.ReadFile("../src/websites/websites.json")
	if err != nil {
		log.Fatal(err)
	}
	var sites []string
	err = json.Unmarshal(file, &sites)
	if err != nil {
		log.Fatal(err)
	}
	return sites
}