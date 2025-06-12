package scraper

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/jaimevidex/go_webscraper/internal/models"
)

func ScrapeArticles(url string) ([]models.Article, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("StatusCode error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		var art models.Article

		// timestamp
		if tsStr, exists := s.Find("span.date").Attr("data-timestamp"); exists {
			if ts, err := strconv.ParseInt(tsStr, 10, 64); err == nil {
				art.Timestamp = ts
			}
		}

		// title and link
		a := s.Find("h3.title a")
		art.Title = a.Find("span").Text()
		if href, exists := a.Attr("href"); exists {
			art.Link = "https://24.sapo.pt" + href
		}

		// description
		art.Description = s.Find("div.excerpt").Text()

		articles = append(articles, art)
	})

	return articles, nil
}
