package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/jaimevidex/go_webscraper/internal/scraper"
)

func main() {
	url := "https://24.sapo.pt/atualidade"
	articles, err := scraper.ScrapeArticles(url)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("")
	if err != nil {
		log.Fatalf("Failed to create CSV file %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write([]string{"Date", "Title", "Link", "Description"})
	if err != nil {
		log.Fatalf("Error wrtiting CSV header: %v", err)
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
			log.Printf("Error writing row: %v", err)
		}
	}

	log.Println("CSV file 'articles.csv' created successfully.")
}
