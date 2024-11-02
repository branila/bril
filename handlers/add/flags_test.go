package adder

import (
	"os"
	"strings"
	"testing"
)

func TestGetFlags(t *testing.T) {
	tests := []struct {
		args []string
		want AddFlags
	}{
		{
			[]string{
				"bril",
				"add",
				"Task name",
				"-p", "5",
				"-d", "2021-12-31",
				"-t", "work",
				"-n", "do the hard things",
			},
			AddFlags{
				Priority: 5,
				Due:      "2021-12-31",
				Tag:      "work",
				Note:     "do the hard things",
			},
		},
		{
			[]string{
				"bril",
				"add",
				"Task name",
			},
			AddFlags{
				Priority: -1,
			},
		},
		{
			[]string{
				"bril",
				"add",
				"Task name",
				"-p", "5",
				"-d", "2021-12-31",
				"-t", "work",
			},
			AddFlags{
				Priority: 5,
				Due:      "2021-12-31",
				Tag:      "work",
			},
		},
		{
			[]string{
				"bril",
				"add",
				"Task name",
				"-p", "5",
				"-d", "2021-12-31",
			},
			AddFlags{
				Priority: 5,
				Due:      "2021-12-31",
			},
		},
		{
			[]string{
				"bril",
				"add",
				"Task name",
				"-p", "5",
				"-n", "do the hard things",
			},
			AddFlags{
				Priority: 5,
				Note:     "do the hard things",
			},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			os.Args = tt.args

			got := getFlags()

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
