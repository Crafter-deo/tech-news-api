package websites

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeMashable() []news {

	doc, err := getMashableHtml()
	if err != nil {
		return nil
	}
	var listOfNews []news

	filterMashabledivs := func(i int, s *goquery.Selection) bool {
		return strings.TrimSpace(s.Find("h2.font-bold.header-200").Text()) == "Latest"
	}

	doc.Find("div.max-w-8xl.px-4.mx-auto.pb-8.mt-12").FilterFunction(filterMashabledivs).Find("div.flex-1").Each(func(i int, s *goquery.Selection) {
		topicCard := news{}
		baseUrl := "https://mashable.com"
		topicCard.Headline = s.Find("div.flex-1 a.block.w-full.text-lg.font-bold.leading-6.mt-2").Text()
		href, _ := s.Find("div.flex-1 a.block.w-full.text-lg.font-bold.leading-6.mt-2").Attr("href")
		topicCard.Url = baseUrl + href
		topicCard.Site = "Mashable"

		listOfNews = append(listOfNews, topicCard)
	})

	return listOfNews
}
func getMashableHtml() (*goquery.Document, error) {
	url := "https://mashable.com/tech"

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
