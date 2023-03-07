package websites

import (
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeCodingdojo(wg *sync.WaitGroup, channel chan<- []Blogs) {
	defer wg.Done()
	doc, err := getCodingdojoHtml()

	if err != nil {
		log.Println("Error parsing document", err)
	}

	var listOfNews []Blogs

	doc.Find("div.jet-smart-listing__post-wrapper").Each(func(i int, s *goquery.Selection) {
		var topicCard Blogs
		topicCard.Headline = s.Find("div.jet-smart-listing__post-title.post-title-simple a").Text()
		topicCard.Url, _ = s.Find("div.jet-smart-listing__post-title.post-title-simple a").Attr("href")
		topicCard.Site = "Codingdojo"

		listOfNews = append(listOfNews, topicCard)
	})
	if len(listOfNews) > 5 {
		channel <- listOfNews[:5]
	} else {
		channel <- listOfNews
	}

}

func getCodingdojoHtml() (*goquery.Document, error) {
	url := "https://www.codingdojo.com/blog/"

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
