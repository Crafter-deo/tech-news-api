package websites

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeCodingdojo() []Blogs {
	doc, err := getCodingdojoHtml()

	if err != nil {
		return nil
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
		return listOfNews[:5]
	} else {
		return listOfNews
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
