package websites

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeCnet() ([]news, error) {
	doc, err := getCnetHtml()
	if err != nil {
		return nil, err
	}
	var listOfNews []news
	baseUrl := "https://www.cnet.com"
	doc.Find("div.o-card.c-premiumCards_card").Each(func(i int, selection *goquery.Selection) {
		topicCard := news{}
		topicCard.Headline = selection.Find("h3.c-premiumCards_title").Text()
		href, _ := selection.Find("a.o-linkOverlay").Attr("href")
		topicCard.Url = baseUrl + href
		topicCard.Site = "Cnet"
		listOfNews = append(listOfNews, topicCard)
	})

	return listOfNews[:5], nil
}

func getCnetHtml() (*goquery.Document, error) {
	client := &http.Client{}
	url := "https://www.cnet.com/tech/"
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
