# go_webscraper

A small, dependency-light web scraper written in **Go**. It walks the news archive of [24.sapo.pt](https://24.sapo.pt) page by page and exports every article to a CSV file.

Built to answer a simple question — *what did the Portuguese news cycle look like over the last decade?* — and as an excuse to write idiomatic Go with a clean package layout.

## What it extracts

| Column | Source |
|---|---|
| `Date` | `data-timestamp` on the article element |
| `Title` | `h3.title a span` |
| `Link` | `h3.title a[href]`, resolved against the site root |
| `Description` | `div.excerpt` |

## Usage

```bash
go run ./cmd/webscraper -pages 50 -out articles.csv
```

| Flag | Default | Meaning |
|---|---|---|
| `-pages` | `9526` | Number of archive pages to walk |
| `-out` | `articles.csv` | Output file |
| `-base` | `https://24.sapo.pt/atualidade?pagina=%d` | URL template (`%d` = page) |

Build a standalone binary with:

```bash
go build -o webscraper ./cmd/webscraper
```

## Layout

```
cmd/webscraper/main.go     CLI: flags, CSV writer, page loop
internal/scraper/          HTTP fetch + goquery parsing
internal/models/           Article type
internal/utils/            Date/timestamp helpers
```

The scraper uses a single shared `http.Client` with a 30-second timeout so connections are reused across thousands of pages, and sets an explicit `User-Agent`. Failed pages are logged and skipped instead of aborting the run.

## Design notes

- **Streaming output.** Rows are flushed to the CSV as they arrive, so a long run can be interrupted without losing everything.
- **No retries.** A page that fails is skipped; the archive is large enough that this does not matter for aggregate analysis.
- **Selector fragility.** The parser is tied to the site's markup. If the layout changes, `internal/scraper/scraper.go` is the only file to touch.

## Before you run it

This is an educational crawler. Respect the site's `robots.txt` and terms of service, keep the request rate low, and do not republish scraped content as your own. The default of 9,526 pages is the full archive — for experiments, start with `-pages 5`.

## Possible next steps

- [ ] Adapt the parser to other Portuguese news sites (SIC Notícias, Público, …)
- [ ] Add a generic, per-site configuration layer
- [ ] Add retry with backoff and a polite rate limit
