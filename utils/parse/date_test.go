package utils

import (
	"testing"
	"time"
)

func TestFindDateFormat(t *testing.T) {
	tests := []struct {
		due        string
		expected   string
		shouldFind bool
	}{
		{"31/12/2024-14:30", "02/01/2006-15:04", true},
		{"14:30-31/12/2024", "15:04-02/01/2006", true},
		{"31/12/2024", "02/01/2006", true},
		{"01-02-2024", "02-01-2006", true},
		{"31/12", "02/01", true},
		{"14:30", "15:04", true},
		{"14.30", "15.04", true},
		{"70/02/2024", "02/01/2006", true},
		{"05/2024", "", false},
		{"14-30", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.due, func(t *testing.T) {
			layout, found := findDateFormat(tt.due)

			if found != tt.shouldFind {
				t.Fatalf("expected found to be %v, got %v", tt.shouldFind, found)
			}

			if found && layout != tt.expected {
				t.Fatalf("expected layout %q, got %q", tt.expected, layout)
			}
		})
	}
}

func TestAdjustTime(t *testing.T) {
	now := time.Now()

	tests := []struct {
		time time.Time
		want time.Time
	}{
		{
			now.Add(-time.Hour),
			now.Add(-time.Hour).AddDate(0, 0, 1),
		},
		{
			now.Add(time.Hour),
			now.Add(time.Hour),
		},
		{
			now.Add(-time.Minute),
			now.Add(-time.Minute).AddDate(0, 0, 1),
		},
		{
			now.Add(time.Hour),
			now.Add(time.Hour),
		},
	}

	for _, tt := range tests {
		t.Run(tt.time.String(), func(t *testing.T) {
			got := AdjustTime(tt.time)
			if !got.Equal(tt.want) {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestAdjustDate(t *testing.T) {
	now := time.Now()

	tests := []struct {
		date time.Time
		want time.Time
	}{
		{
			now.AddDate(0, 0, -1),
			now.AddDate(0, 0, -1).AddDate(1, 0, 0),
		},
		{
			now.AddDate(0, 0, -7),
			now.AddDate(0, 0, -7).AddDate(1, 0, 0),
		},
		{
			now.AddDate(0, 1, 0),
			now.AddDate(0, 1, 0),
		},
		{
			now.AddDate(0, 0, 2),
			now.AddDate(0, 0, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.date.String(), func(t *testing.T) {
			got := AdjustDate(tt.date)
			if !got.Equal(tt.want) {
				t.Fatalf("expected %v, got %v", tt.want, got)
			}
		})
	}
}
