package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/jaimevidex/go_webscraper/internal/scraper"
)

func main() {
	const totalPages = 9526
	baseURL := "https://24.sapo.pt/atualidade?pagina=%d"

	file, err := os.Create("articles.csv")
	if err != nil {
		log.Fatalf("Failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"Date", "Title", "Link", "Description"})
	if err != nil {
		log.Fatalf("Error writing CSV header: %v", err)
	}

	for page := 1; page <= totalPages; page++ {
		url := fmt.Sprintf(baseURL, page)
		log.Printf("Scraping page %d of %d: %s", page, totalPages, url)

		articles, err := scraper.ScrapeArticles(url)
		if err != nil {
			log.Printf("Error scraping page %d: %v", page, err)
			continue // Skip this page
		}

		for _, a := range articles {
			row := []string{
				a.Date().Format("2006-01-02 15:04:05"),
				a.Title,
				a.Link,
				a.Description,
			}
			err := writer.Write(row)
			if err != nil {
				log.Printf("Error writing row from page %d: %v", page, err)
			}
		}
	}

	log.Println("✅ All pages scraped. CSV file 'articles.csv' created successfully.")
}
