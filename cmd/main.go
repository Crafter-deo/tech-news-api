package main

import (
	// "log"
	"net/http"

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
	router.Run()
}
