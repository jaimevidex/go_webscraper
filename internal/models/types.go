package models

import (
	"time"

	"github.com/jaimevidex/go_webscraper/internal/utils"
)

type Article struct {
	Timestamp   int64
	Title       string
	Link        string
	Description string
}

func (a *Article) Date() time.Time {
	return utils.TimestampMillisToTime(a.Timestamp)
}
