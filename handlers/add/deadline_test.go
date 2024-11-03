package adder

import (
	"math"
	"testing"
	"time"

	utils "github.com/branila/bril/utils/parse"
)

func TestGetDeadline(t *testing.T) {
	now := time.Now()

	tests := []struct {
		due         string
		expectedDue time.Time
		expectedErr string
	}{
		{"", time.Time{}, ""},
		{"31/12/2024-14:30", time.Date(2024, 12, 31, 14, 30, 0, 0, time.Local), ""},
		{"14:30-01/02/2024", time.Date(2024, 02, 01, 14, 30, 0, 0, time.Local), ""},
		{"31/12/2024", time.Date(2024, 12, 31, 0, 0, 0, 0, time.Local), ""},
		{"01-02-2024", time.Date(2024, 02, 01, 0, 0, 0, 0, time.Local), ""},
		{"31/12", utils.AdjustDate(time.Date(now.Year(), 12, 31, 0, 0, 0, 0, time.Local)), ""},
		{"14:30", utils.AdjustTime(time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 0, 0, time.Local)), ""},
		{"14.30", utils.AdjustTime(time.Date(now.Year(), now.Month(), now.Day(), 14, 30, 0, 0, time.Local)), ""},
		{"30/02/2024", time.Time{}, "Failed to parse date"},
		{"05/2024", time.Time{}, "Unknown date format"},
		{"14-30", time.Time{}, "Unknown date format"},
		{"1y", now.AddDate(1, 0, 0), ""},
		{"1M", now.AddDate(0, 1, 0), ""},
		{"1d", now.AddDate(0, 0, 1), ""},
		{"1h", now.Add(1 * time.Hour), ""},
		{"1m", now.Add(time.Minute), ""},
		{"1s", now.Add(time.Second), ""},
		{"1y1M1d1h1m1s", now.AddDate(1, 1, 1).Add(time.Hour).Add(time.Minute).Add(time.Second), ""},
		{"1y5w3d", now.AddDate(1, 0, 38), ""},
	}

	for _, tt := range tests {
		t.Run(tt.due, func(t *testing.T) {
			deadline, err := getDeadline(tt.due)

			if err != nil && tt.expectedErr != err.Error() {
				t.Fatalf("expected error %q, got %q", tt.expectedErr, err.Error())
			}

			if math.Abs(float64(deadline.Sub(tt.expectedDue).Milliseconds())) > 1 {
				t.Fatalf("expected deadline to be %v, got %v", tt.expectedDue, deadline)
			}
		})
	}
}
