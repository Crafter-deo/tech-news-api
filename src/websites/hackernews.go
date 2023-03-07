package websites

import (
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeHackernews(wg *sync.WaitGroup, channel chan<- []Blogs) {
	defer wg.Done()
	doc, err := getHackernewsHtml()
	if err != nil {
		log.Println("Error parsing document", err)
	}

	var listOfNews []Blogs

	doc.Find("span.titleline").Each(func(i int, s *goquery.Selection) {
		topicCard := Blogs{}

		topicCard.Headline = s.Find("a").Text()
		topicCard.Url, _ = s.Find("a").Attr("href")
		topicCard.Site = "Hacker News"

		listOfNews = append(listOfNews, topicCard)
	})

	channel <- listOfNews[:5]
}

func getHackernewsHtml() (*goquery.Document, error) {
	url := "https://news.ycombinator.com/"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")

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
