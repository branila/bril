package modifier

import (
	"os"
	"testing"
)

func TestGetFlags(t *testing.T) {
	tests := []struct {
		args []string
		want EditFlags
	}{
		{
			[]string{
				"bril",
				"edit",
				"1",
				"-n", "ascoltare nitro",
				"-s", "done",
				"-t", "work",
				"-p", "5",
				"-note", "do the hard things",
				"-d", "2021-12-31",
			},
			EditFlags{
				Name:     "ascoltare nitro",
				State:    "done",
				Tag:      "work",
				Priority: 5,
				Note:     "do the hard things",
				Deadline: "2021-12-31",
			},
		},
		{
			[]string{
				"bril",
				"edit",
				"2",
				"-n", "ascoltare rhio",
				"-p", "10",
			},
			EditFlags{
				Name:     "ascoltare rhio",
				Priority: 10,
			},
		},
		{
			[]string{
				"bril",
				"edit",
				"3",
				"-s", "deleted",
			},
			EditFlags{
				State:    "deleted",
				Priority: -1,
			},
		},
		{
			[]string{
				"bril",
				"edit",
				"4",
				"-t", "work",
			},
			EditFlags{
				Tag:      "work",
				Priority: -1,
			},
		},
		{
			[]string{
				"bril",
				"edit",
				"5",
				"-p", "5",
			},
			EditFlags{
				Priority: 5,
			},
		},
		{
			[]string{
				"bril",
				"edit",
				"6",
				"-note", "do the hard things",
			},
			EditFlags{
				Note:     "do the hard things",
				Priority: -1,
			},
		},
		{
			[]string{
				"bril",
				"edit",
				"7",
				"-d", "2021-12-31",
			},
			EditFlags{
				Deadline: "2021-12-31",
				Priority: -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.args[2], func(t *testing.T) {
			os.Args = tt.args

			got := getFlags()

			if got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
