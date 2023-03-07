package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sudodeo/tech-news-api/src/websites"
)

const PORT = ":2222"

func main() {
	// log.Println(websites.ScrapeCnet())
	router := gin.Default()
	wg := &sync.WaitGroup{}

	// Create channels for each website
	mediumChannel := make(chan []websites.Blogs)
	mashableChannel := make(chan []websites.Blogs)
	hackernewsChannel := make(chan []websites.Blogs)
	digitaltrendsChannel := make(chan []websites.Blogs)
	codingdojoChannel := make(chan []websites.Blogs)
	cnetChannel := make(chan []websites.Blogs)

	router.GET("/all", func(ctx *gin.Context) {

		// Use a buffered channel to collect all data
		all_blogs := make(chan []websites.Blogs, 6)
		wg.Add(6)
		go websites.ScrapeMedium(wg, mediumChannel)
		go websites.ScrapeMashable(wg, mashableChannel)
		go websites.ScrapeHackernews(wg, hackernewsChannel)
		go websites.ScrapeDigitaltrends(wg, digitaltrendsChannel)
		go websites.ScrapeCodingdojo(wg, codingdojoChannel)
		go websites.ScrapeCnet(wg, cnetChannel)
		wg.Add(1)
		// Fan-in the scraped data into a single channel
		go func() {
			defer wg.Done()
			for i := 0; i < 6; i++ {
				select {
				case data := <-mediumChannel:
					all_blogs <- data
				case data := <-mashableChannel:
					all_blogs <- data
				case data := <-hackernewsChannel:
					all_blogs <- data
				case data := <-digitaltrendsChannel:
					all_blogs <- data
				case data := <-codingdojoChannel:
					all_blogs <- data
				case data := <-cnetChannel:
					all_blogs <- data
				}
			}
			close(all_blogs)
		}()

		// Collect all data from the channel
		allData := []websites.Blogs{}
		for data := range all_blogs {
			allData = append(allData, data...)
		}
		wg.Wait()

		// Return all data as JSON
		ctx.JSON(http.StatusOK, allData)
	})

	router.GET("/cnet", func(ctx *gin.Context) {
		wg.Add(1)
		go websites.ScrapeCnet(wg, cnetChannel)
		ctx.JSON(http.StatusOK, <-cnetChannel)
	})
	router.GET("/codingdojo", func(ctx *gin.Context) {
		wg.Add(1)
		go websites.ScrapeCodingdojo(wg, codingdojoChannel)
		ctx.JSON(http.StatusOK, <-codingdojoChannel)
	})
	router.GET("/digitaltrends", func(ctx *gin.Context) {
		wg.Add(1)
		go websites.ScrapeDigitaltrends(wg, digitaltrendsChannel)
		ctx.JSON(http.StatusOK, <-digitaltrendsChannel)
	})
	router.GET("/hackernews", func(ctx *gin.Context) {
		wg.Add(1)
		go websites.ScrapeHackernews(wg, hackernewsChannel)
		ctx.JSON(http.StatusOK, <-hackernewsChannel)
	})
	router.GET("/mashable", func(ctx *gin.Context) {
		wg.Add(1)
		go websites.ScrapeMashable(wg, mashableChannel)
		ctx.JSON(http.StatusOK, <-mashableChannel)

	})
	router.GET("/medium", func(ctx *gin.Context) {
		wg.Add(1)
		go websites.ScrapeMedium(wg, mediumChannel)
		ctx.JSON(http.StatusOK, <-mediumChannel)
	})

	router.Run(PORT)
}
