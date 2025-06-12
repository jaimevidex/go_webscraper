package main

import (
	"fmt"
	"github.com/jaimevidex/go_webscraper/internal/scraper"
	"log"
)

func main() {
	url := "https://24.sapo.pt/atualidade"
	articles, err := scraper.ScrapeArticles(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, a := range articles {
		fmt.Printf("Timestamp: %d\nTitle: %s\nLink: %s\nDescription: %s\n\n",
			a.Timestamp, a.Title, a.Link, a.Description)
	}
}

