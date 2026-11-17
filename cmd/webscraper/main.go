package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jaimevidex/go_webscraper/internal/scraper"
)

func main() {
	pages := flag.Int("pages", 9526, "number of archive pages to scrape")
	out := flag.String("out", "articles.csv", "path of the CSV file to write")
	base := flag.String("base", "https://24.sapo.pt/atualidade?pagina=%d", "archive URL template with %d for the page number")
	flag.Parse()

	file, err := os.Create(*out)
	if err != nil {
		log.Fatalf("failed to create %s: %v", *out, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{"Date", "Title", "Link", "Description"}); err != nil {
		log.Fatalf("failed to write CSV header: %v", err)
	}

	articles := 0
	for page := 1; page <= *pages; page++ {
		url := fmt.Sprintf(*base, page)
		log.Printf("scraping page %d/%d", page, *pages)

		found, err := scraper.ScrapeArticles(url)
		if err != nil {
			log.Printf("page %d failed: %v", page, err)
			continue
		}

		for _, a := range found {
			row := []string{
				a.Date().Format("2006-01-02 15:04:05"),
				a.Title,
				a.Link,
				a.Description,
			}
			if err := writer.Write(row); err != nil {
				log.Printf("page %d: failed to write row: %v", page, err)
			}
		}
		articles += len(found)
	}

	log.Printf("done: %d articles written to %s", articles, *out)
}
