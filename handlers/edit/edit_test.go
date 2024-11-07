package modifier

import (
	"testing"
	"time"

	"github.com/branila/bril/types"
	utils "github.com/branila/bril/utils/parse"
)

func TestAllFlagsUnset(t *testing.T) {
	tests := []struct {
		name string
		args EditFlags
		want bool
	}{
		{
			"all flags unset",
			EditFlags{
				Priority: -1,
			},
			true,
		},
		{
			"a flag set",
			EditFlags{
				Priority: -1,
				Name:     "ascoltare i punkreas",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allFlagsUnset(tt.args); got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEditedTask(t *testing.T) {
	testDate, _ := utils.ParseDate("13-11-2024")

	tests := []struct {
		name  string
		task  types.Task
		flags EditFlags
		want  types.Task
		error string
	}{
		{
			"edit name",
			types.Task{
				Id:   1,
				Name: "ascoltare i punkreas",
			},
			EditFlags{
				Name: "ascoltare rhio",
			},
			types.Task{
				Id:   1,
				Name: "ascoltare rhio",
			},
			"",
		},
		{
			"edit priority",
			types.Task{
				Id:       1,
				Name:     "ascoltare i punkreas",
				Priority: 5,
			},
			EditFlags{
				Priority: 3,
			},
			types.Task{
				Id:       1,
				Name:     "ascoltare i punkreas",
				Priority: 3,
			},
			"",
		},
		{
			"edit deadline",
			types.Task{
				Id:       1,
				Name:     "ascoltare i punkreas",
				Deadline: time.Now(),
			},
			EditFlags{
				Deadline: "13-11-2024",
			},
			types.Task{
				Id:       1,
				Name:     "ascoltare i punkreas",
				Deadline: testDate,
			},
			"",
		},
		{
			"edit state",
			types.Task{
				Id:   1,
				Name: "ascoltare i punkreas",
				Done: false,
			},
			EditFlags{
				State: "done",
			},
			types.Task{
				Id:   1,
				Name: "ascoltare i punkreas",
				Done: true,
			},
			"",
		},
		{
			"edit with invalid state",
			types.Task{
				Id:   1,
				Name: "ascoltare i punkreas",
				Done: false,
			},
			EditFlags{
				State: "fioritura",
			},
			types.Task{},
			"Invalid state",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getEditedTask(tt.task, tt.flags)

			if err != nil && err.Error() != tt.error {
				t.Errorf("Got error %v, want %v", err.Error(), tt.error)
			}

			if got != tt.want {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}
