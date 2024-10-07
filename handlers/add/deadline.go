package handlers

import (
	"log"
	"regexp"
	"time"
)

func getDeadline(due string) time.Time {
	if due == "" {
		return time.Time{}
	}

	layout, found := findDateFormat(due)
	if !found {
		log.Fatal("Unknown date format")
	}

	deadline, err := time.Parse(layout, due)
	if err != nil {
		log.Fatal(err)
	}

	if layout == "15:04" {
		deadline = adjustTime(deadline)
	}

	return deadline
}

func findDateFormat(due string) (string, bool) {
	formats := map[string]string{
		`^\d{2}/\d{2}/\d{4}-\d{2}:\d{2}$`: "02/01/2006-15:04",
		`^\d{2}:\d{2}-\d{2}/\d{2}/\d{4}$`: "15:04-02/01/2006",
		`^\d{2}/\d{2}/\d{4}$`:             "02/01/2006",
		`^\d{2}:\d{2}$`:                   "15:04",
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
		t.Hour(), t.Minute(), 0, 0,
		time.Local,
	)

	if t.Before(now) {
		t = t.Add(24 * time.Hour)
	}

	return t
}
