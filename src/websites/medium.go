package websites

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeMedium() []Blogs {
	doc, err := getMediumHtml()
	if err != nil {
		return nil
	}

	var listOfNews []Blogs
	doc.Find("div.kp.l").Each(func(i int, s *goquery.Selection) {
		var topicCard Blogs
		baseUrl := "https://medium.com"
		href, _ := s.Find(`a[aria-label="Post Preview Title"]`).Attr("href")
		topicCard.Url = baseUrl + href
		topicCard.Site = "Medium"
		topicCard.Headline = s.Find("h2").Text()
		listOfNews = append(listOfNews, topicCard)

	})

	return listOfNews[:5]

}

func getMediumHtml() (*goquery.Document, error) {
	url := "https://medium.com/tag/technology"

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
