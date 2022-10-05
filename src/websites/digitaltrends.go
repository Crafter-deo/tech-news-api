package websites

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeDigitaltrends() []Blogs {
	doc, err := getDigitaltrendsHtml()
	if err != nil {
		return nil
	}

	listOfNews := []Blogs{}

	doc.Find("div.b-tabbed-lists__items.h-tabbed-lists-list.is-active div.b-tabbed-lists__item").Each(func(i int, s *goquery.Selection) {
		topicCard := Blogs{}

		topicCard.Headline = s.Find("h3 a").Text()
		topicCard.Url, _ = s.Find("h3 a").Attr("href")
		topicCard.Site = "Digital Trends"

		listOfNews = append(listOfNews, topicCard)

	})

	return listOfNews
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
