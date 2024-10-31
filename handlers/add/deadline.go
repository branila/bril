package adder

import (
	"fmt"
	"regexp"
	"time"
)

func getDeadline(due string) (time.Time, error) {
	if due == "" {
		return time.Time{}, nil
	}

	layout, found := findDateFormat(due)
	if !found {
		return time.Time{}, fmt.Errorf("Unknown date format")
	}

	deadline, err := time.ParseInLocation(layout, due, time.Local)
	if err != nil {
		return time.Time{}, fmt.Errorf("Failed to parse date")
	}

	if layout == "15:04" || layout == "15.04" {
		deadline = adjustTime(deadline)
	}

	if layout == "02/01" || layout == "02-01" {
		deadline = adjustDate(deadline)
	}

	return deadline, nil
}

func findDateFormat(due string) (string, bool) {
	formats := map[string]string{
		`^\d{2}/\d{2}/\d{4}-\d{2}:\d{2}$`: "02/01/2006-15:04",
		`^\d{2}:\d{2}-\d{2}/\d{2}/\d{4}$`: "15:04-02/01/2006",
		`^\d{2}/\d{2}/\d{4}$`:             "02/01/2006",
		`^\d{2}-\d{2}-\d{4}$`:             "02-01-2006",
		`^\d{2}/\d{2}$`:                   "02/01",
		`^\d{2}:\d{2}$`:                   "15:04",
		`^\d{2}\.\d{2}$`:                  "15.04",
	}

	for pattern, layout := range formats {
		if matched, _ := regexp.MatchString(pattern, due); matched {
			return layout, true
		}
	}

	return "", false
}

func adjustTime(t time.Time) time.Time {
	now := time.Now()

	t = time.Date(
		now.Year(), now.Month(), now.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
		time.Local,
	)

	if t.Before(now) {
		t = t.Add(24 * time.Hour)
	}

	return t
}

func adjustDate(t time.Time) time.Time {
	now := time.Now()

	t = time.Date(
		now.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
		time.Local,
	)

	if t.Before(now) {
		t = t.AddDate(1, 0, 0)
	}

	return t
}
