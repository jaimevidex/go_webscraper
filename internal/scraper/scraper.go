package scraper

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jaimevidex/go_webscraper/internal/models"
)

// client is shared so connections are reused across pages.
var client = &http.Client{Timeout: 30 * time.Second}

// ScrapeArticles extracts every article from a single archive page.
func ScrapeArticles(url string) ([]models.Article, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; go_webscraper/1.0)")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var articles []models.Article

	doc.Find("article").Each(func(_ int, s *goquery.Selection) {
		var art models.Article

		if tsStr, exists := s.Find("span.date").Attr("data-timestamp"); exists {
			if ts, err := strconv.ParseInt(tsStr, 10, 64); err == nil {
				art.Timestamp = ts
			}
		}

		link := s.Find("h3.title a")
		art.Title = link.Find("span").Text()
		if href, exists := link.Attr("href"); exists {
			art.Link = "https://24.sapo.pt" + href
		}

		art.Description = s.Find("div.excerpt").Text()

		articles = append(articles, art)
	})

	return articles, nil
}
