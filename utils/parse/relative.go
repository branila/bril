package utils

import (
	"fmt"
	"strings"
	"time"
)

func IsRelativeFormat(inputDeadline string) bool {
	validFormats := []string{
		"today", "tomorrow",
		"next week", "next month", "next year",
		"monday", "next monday",
		"tuesday", "next tuesday",
		"wednesday", "next wednesday",
		"thursday", "next thursday",
		"friday", "next friday",
		"saturday", "next saturday",
		"sunday", "next sunday",
	}

	for _, format := range validFormats {
		if inputDeadline == format {
			return true
		}
	}

	return false
}

func ParseRelativeFormat(inputDeadline string) (time.Time, error) {
	inputDeadline = strings.ToLower(inputDeadline)

	now := time.Now().Truncate(24 * time.Hour)

	relativeTime := map[string]func() time.Time{
		"today":     func() time.Time { return now.AddDate(0, 0, 1) },
		"tomorrow":  func() time.Time { return now.AddDate(0, 0, 2) },
		"next week": func() time.Time { return now.AddDate(0, 0, int(time.Sunday-now.Weekday())+7).Truncate(24 * time.Hour) },
		"next month": func() time.Time {
			return time.Date(now.Year(), now.Month()+2, 1, 0, 0, 0, 0, now.Location()).AddDate(0, 0, -1)
		},
		"next year": func() time.Time { return time.Date(now.Year()+2, 1, 1, 0, 0, 0, 0, now.Location()).AddDate(0, 0, -1) },
	}

	weekdays := []time.Weekday{
		time.Monday, time.Tuesday, time.Wednesday, time.Thursday,
		time.Friday, time.Saturday, time.Sunday,
	}

	for _, day := range weekdays {
		dayStr := strings.ToLower(day.String())

		relativeTime[dayStr] = func(d time.Weekday) func() time.Time {
			return func() time.Time { return getWeekday(d) }
		}(day)

		relativeTime["next "+dayStr] = func(d time.Weekday) func() time.Time {
			return func() time.Time { return getWeekday(d).AddDate(0, 0, 7) }
		}(day)
	}

	if fn, exists := relativeTime[inputDeadline]; exists {
		return fn(), nil
	}

	return time.Time{}, fmt.Errorf("Invalid relative format")
}

func getWeekday(weekday time.Weekday) time.Time {
	now := time.Now().Truncate(24 * time.Hour)
	daysUntil := int(weekday - now.Weekday())

	if daysUntil <= 0 {
		daysUntil += 7
	}

	return now.AddDate(0, 0, daysUntil)
}
