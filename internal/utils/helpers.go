package utils

import (
	"time"
)

// TimestampMillisToTime converts a Unix timestamp in milliseconds to time.Time
func TimestampMillisToTime(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Millisecond))
}
