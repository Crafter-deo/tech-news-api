package websites

import (
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func ScrapeDigitaltrends(ctx *gin.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	doc, err := getDigitaltrendsHtml()
	if err != nil {
		log.Println("Error parsing document", err)
	}

	listOfNews := []Blogs{}

	doc.Find("div.b-mem-post__content").Each(func(i int, s *goquery.Selection) {
		topicCard := Blogs{}

		topicCard.Headline = s.Find("h3 a").Text()
		topicCard.Url, _ = s.Find("h3 a").Attr("href")
		topicCard.Site = "Digital Trends"

		listOfNews = append(listOfNews, topicCard)

	})

	if len(listOfNews) > 5 {
		ctx.JSON(http.StatusOK, listOfNews[:5])
	} else {
		ctx.JSON(http.StatusOK, listOfNews)
	}
}

func getDigitaltrendsHtml() (*goquery.Document, error) {
	url := "https://www.digitaltrends.com/news/"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
