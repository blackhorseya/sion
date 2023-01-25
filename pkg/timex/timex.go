package timex

import (
	"time"
)

const (
	// RFC3339Mill RFC3339 format with millisecond
	RFC3339Mill = "2006-01-02T15:04:05.000Z07:00"
)

// ParseYYYYMMddHHmmss serve caller to given string to parse time
func ParseYYYYMMddHHmmss(str string) (time.Time, error) {
	loc, _ := time.LoadLocation("Asia/Taipei")
	layout := "20060102150405"

	t, err := time.ParseInLocation(layout, str, loc)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
