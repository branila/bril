package utils

import (
	"testing"
	"time"
)

func TestIsRelativeFormat(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"today", true},
		{"tomorrow", true},
		{"next week", true},
		{"next month", true},
		{"next year", true},
		{"monday", true},
		{"next monday", true},
		{"tuesday", true},
		{"next tuesday", true},
		{"wednesday", true},
		{"next wednesday", true},
		{"thursday", true},
		{"next thursday", true},
		{"friday", true},
		{"next friday", true},
		{"saturday", true},
		{"next saturday", true},
		{"sunday", true},
		{"next sunday", true},
		// Negative cases
		{"yesterday", false},
		{"next year 2024", false},
		{"weekend", false},
		{"next month end", false},
		{"invalid format", false},
		{"", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := IsRelativeFormat(test.input)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestParseRelativeFormat(t *testing.T) {
	now := time.Now().Truncate(24 * time.Hour)

	tests := []struct {
		input         string
		expected      time.Time
		expectedError string
	}{
		{"today", now.AddDate(0, 0, 1), ""},
		{"tomorrow", now.AddDate(0, 0, 2), ""},
		{"next week", now.AddDate(0, 0, int(time.Sunday-now.Weekday())+7), ""},
		{"next month", time.Date(now.Year(), now.Month()+2, 1, 0, 0, 0, 0, now.Location()).AddDate(0, 0, -1), ""},
		{"next year", time.Date(now.Year()+2, 1, 1, 0, 0, 0, 0, now.Location()).AddDate(0, 0, -1), ""},
		{"monday", getWeekday(time.Monday), ""},
		{"next monday", getWeekday(time.Monday).AddDate(0, 0, 7), ""},
		{"tuesday", getWeekday(time.Tuesday), ""},
		{"next tuesday", getWeekday(time.Tuesday).AddDate(0, 0, 7), ""},
		{"wednesday", getWeekday(time.Wednesday), ""},
		{"next wednesday", getWeekday(time.Wednesday).AddDate(0, 0, 7), ""},
		{"thursday", getWeekday(time.Thursday), ""},
		{"next thursday", getWeekday(time.Thursday).AddDate(0, 0, 7), ""},
		{"friday", getWeekday(time.Friday), ""},
		{"next friday", getWeekday(time.Friday).AddDate(0, 0, 7), ""},
		{"saturday", getWeekday(time.Saturday), ""},
		{"next saturday", getWeekday(time.Saturday).AddDate(0, 0, 7), ""},
		{"sunday", getWeekday(time.Sunday), ""},
		{"next sunday", getWeekday(time.Sunday).AddDate(0, 0, 7), ""},
		// Negative cases
		{"yesterday", time.Time{}, "Invalid relative format"},
		{"next year 2024", time.Time{}, "Invalid relative format"},
		{"weekend", time.Time{}, "Invalid relative format"},
		{"next month end", time.Time{}, "Invalid relative format"},
		{"invalid format", time.Time{}, "Invalid relative format"},
		{"", time.Time{}, "Invalid relative format"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := ParseRelativeFormat(test.input)

			if err != nil && err.Error() != test.expectedError {
				t.Errorf("expected error message: %v, got: %v", test.expectedError, err.Error())
			}

			if err == nil && test.expectedError != "" {
				t.Errorf("expected error message: %v, but got none", test.expectedError)
			}

			if test.expectedError == "" && !result.Equal(test.expected) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestGetWeekday(t *testing.T) {
	now := time.Now().Truncate(24 * time.Hour)

	weekdays := []time.Weekday{
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
		time.Sunday,
	}

	for _, weekday := range weekdays {
		t.Run(weekday.String(), func(t *testing.T) {
			daysUntil := int(weekday - now.Weekday())

			if daysUntil <= 0 {
				daysUntil += 7
			}

			want := now.AddDate(0, 0, daysUntil)

			got := getWeekday(weekday).Truncate(24 * time.Hour)

			if !got.Equal(want) {
				t.Errorf("Got %v; want %v", got, want)
			}
		})
	}
}
