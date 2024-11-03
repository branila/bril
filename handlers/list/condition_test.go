package lister

import (
	"fmt"
	"testing"
	"time"

	"github.com/branila/bril/types"
)

func TestShouldPrint(t *testing.T) {
	now := time.Now()

	tests := []struct {
		task  types.Task
		flags ListFlags
		want  bool
	}{
		{
			types.Task{
				Id: 0,
			},
			ListFlags{},
			true,
		},
		{
			types.Task{
				Id:   1,
				Done: true,
			},
			ListFlags{
				Done: true,
			},
			true,
		},
		{
			types.Task{
				Id:      2,
				Deleted: true,
			},
			ListFlags{
				Deleted: true,
			},
			true,
		},
		{
			types.Task{
				Id:       3,
				Deadline: now.Add(-time.Hour),
			},
			ListFlags{
				Expired: true,
			},
			true,
		},
		{
			types.Task{
				Id:  4,
				Tag: "work",
			},
			ListFlags{
				Tag: "work",
			},
			true,
		},
		{
			types.Task{
				Id:   5,
				Done: true,
			},
			ListFlags{
				Deleted: true,
			},
			false,
		},
		{
			types.Task{
				Id:      6,
				Deleted: true,
			},
			ListFlags{
				Done: true,
			},
			false,
		},
		{
			types.Task{
				Id:       7,
				Deadline: time.Now().AddDate(10, 0, 0),
			},
			ListFlags{
				Expired: true,
			},
			false,
		},
		{
			types.Task{
				Id:  8,
				Tag: "work",
			},
			ListFlags{
				Tag: "home",
			},
			false,
		},
		{
			types.Task{
				Id:      9,
				Done:    true,
				Deleted: true,
			},
			ListFlags{
				Done:    true,
				Deleted: true,
			},
			true,
		},
		{
			types.Task{
				Id:      10,
				Done:    false,
				Deleted: true,
			},
			ListFlags{
				Deleted: true,
				Done:    true,
			},
			false,
		},
		{
			types.Task{
				Id:       11,
				Done:     true,
				Deadline: now.Add(-time.Hour),
			},
			ListFlags{
				Expired: true,
				Done:    true,
			},
			true,
		},
		{
			types.Task{
				Id:       12,
				Deleted:  true,
				Deadline: now.Add(-time.Hour),
			},
			ListFlags{
				Expired: true,
				Deleted: true,
			},
			true,
		},
		{
			types.Task{
				Id:       13,
				Done:     true,
				Deleted:  true,
				Deadline: now.Add(-time.Hour),
			},
			ListFlags{
				Expired: true,
				Done:    true,
				Deleted: true,
			},
			true,
		},
		{
			types.Task{
				Id:   14,
				Done: true,
			},
			ListFlags{
				Done:    true,
				Deleted: true,
			},
			false,
		},
		{
			types.Task{
				Id:       15,
				Deleted:  true,
				Deadline: now.Add(time.Hour),
			},
			ListFlags{
				Expired: true,
				Deleted: true,
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Id:%d", tt.task.Id), func(t *testing.T) {
			got := shouldPrint(tt.task, tt.flags)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
