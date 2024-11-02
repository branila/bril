package lister

import (
	"os"
	"strings"
	"testing"
)

func TestGetFlags(t *testing.T) {
	tests := []struct {
		args []string
		want ListFlags
	}{
		{
			[]string{
				"bril",
				"list",
				"-a",
			},
			ListFlags{
				All: true,
			},
		},
		{
			[]string{
				"bril",
				"list",
				"-d",
			},
			ListFlags{
				Done: true,
			},
		},
		{
			[]string{
				"bril",
				"list",
				"-t", "work",
			},
			ListFlags{
				Tag: "work",
			},
		},
		{
			[]string{
				"bril",
				"list",
				"-e",
			},
			ListFlags{
				Expired: true,
			},
		},
		{
			[]string{
				"bril",
				"list",
				"-del",
			},
			ListFlags{
				Deleted: true,
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
