package utils

import (
	"math"
	"testing"
	"time"
)

func TestIsDurationFormat(t *testing.T) {
	tests := []struct {
		due        string
		isDuration bool
	}{
		{"", false},
		{"31/12/2024-14:30", false},
		{"14:30-31/12/2024", false},
		{"31/12/2024", false},
		{"01-02-2024", false},
		{"31/12", false},
		{"14:30", false},
		{"14.30", false},
		{"70/02/2024", false},
		{"05/2024", false},
		{"14-30", false},
		{"1y", true},
		{"1M", true},
		{"1d", true},
		{"1h", true},
		{"1m", true},
		{"1s", true},
		{"1y1M1d1h1m1s", true},
		{"3y3M3d3h3m3s", true},
		{"1y1M", true},
		{"1y5w3d", true},
		{"1h1h", true},
		{"1m1m", true},
		{"1m3M2y", true},
		{"3mm", false},
		{"3Mm", false},
		{"h3y", false},
		{"3y3", false},
	}

	for _, tt := range tests {
		t.Run(tt.due, func(t *testing.T) {
			if IsDurationFormat(tt.due) != tt.isDuration {
				t.Fatalf("expected %t, got %t", tt.isDuration, IsDurationFormat(tt.due))
			}
		})
	}
}

func TestParseDuration(t *testing.T) {
	tests := []struct {
		due           string
		expectedDate  time.Time
		expectedError string
	}{
		{"1y", time.Now().AddDate(1, 0, 0), ""},
		{"1M", time.Now().AddDate(0, 1, 0), ""},
		{"1d", time.Now().AddDate(0, 0, 1), ""},
		{"1h", time.Now().Add(time.Hour), ""},
		{"1m", time.Now().Add(time.Minute), ""},
		{"1s", time.Now().Add(time.Second), ""},
		{"1y1M1d1h1m1s", time.Now().AddDate(1, 1, 1).Add(time.Hour).Add(time.Minute).Add(time.Second), ""},
		{"1y1M", time.Now().AddDate(1, 1, 0), ""},
		{"1M1d", time.Now().AddDate(0, 1, 1), ""},
		{"1d1h", time.Now().AddDate(0, 0, 1).Add(time.Hour), ""},
		{"1y5w3d", time.Now().AddDate(1, 0, 38), ""},
		{"1h1h", time.Now().Add(2 * time.Hour), ""},
		{"1h1m", time.Now().Add(time.Hour).Add(time.Minute), ""},
		{"1s2M3y", time.Now().AddDate(3, 2, 0).Add(time.Second), ""},
	}

	for _, tt := range tests {
		t.Run(tt.due, func(t *testing.T) {
			deadline, err := ParseDuration(tt.due)

			if err != nil && tt.expectedError != err.Error() {
				t.Fatalf("expected error %q, got %q", tt.expectedError, err.Error())
			}

			if math.Abs(float64(deadline.Sub(tt.expectedDate).Milliseconds())) > 1 {
				t.Fatalf("expected deadline to be %v, got %v", tt.expectedDate, deadline)
			}
		})
	}
}
