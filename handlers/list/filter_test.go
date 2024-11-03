package lister

import (
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/branila/bril/db"
	"github.com/branila/bril/types"
)

func TestGetFilteredTasks(t *testing.T) {
	prototypeTasks := []types.Task{
		{
			Id:       1,
			Done:     false,
			Deleted:  false,
			Deadline: time.Now().Add(-time.Hour),
			Tag:      "work",
		},
		{
			Id:       2,
			Done:     true,
			Deleted:  false,
			Deadline: time.Now().Add(-24 * time.Hour),
			Tag:      "work",
		},
		{
			Id:       3,
			Done:     false,
			Deleted:  true,
			Deadline: time.Now().Add(time.Hour),
			Tag:      "school",
		},
		{
			Id:       4,
			Done:     true,
			Deleted:  false,
			Deadline: time.Now().Add(24 * time.Hour),
			Tag:      "school",
		},
		{
			Id:       5,
			Done:     false,
			Deleted:  false,
			Deadline: time.Now().AddDate(1, 0, 0),
			Tag:      "work",
		},
		{
			Id:       6,
			Done:     true,
			Deleted:  true,
			Deadline: time.Now().Add(-24 * time.Hour * 7),
			Tag:      "work",
		},
	}

	db.SetTasks(prototypeTasks)

	tests := []struct {
		flags ListFlags
		want  []types.Task
	}{
		{
			ListFlags{},
			[]types.Task{
				prototypeTasks[0],
				prototypeTasks[4],
			},
		},
		{
			ListFlags{
				Done: true,
			},
			[]types.Task{
				prototypeTasks[1],
				prototypeTasks[3],
				prototypeTasks[5],
			},
		},
		{
			ListFlags{
				Deleted: true,
			},
			[]types.Task{
				prototypeTasks[2],
				prototypeTasks[5],
			},
		},
		{
			ListFlags{
				Expired: true,
			},
			[]types.Task{
				prototypeTasks[0],
				prototypeTasks[1],
				prototypeTasks[5],
			},
		},
		{
			ListFlags{
				Tag: "work",
			},
			[]types.Task{
				prototypeTasks[0],
				prototypeTasks[1],
				prototypeTasks[4],
				prototypeTasks[5],
			},
		},
		{
			ListFlags{
				Tag: "school",
			},
			[]types.Task{
				prototypeTasks[2],
				prototypeTasks[3],
			},
		},
		{
			ListFlags{
				All: true,
			},
			prototypeTasks,
		},
		{
			ListFlags{
				Deleted: true,
				Expired: true,
			},
			[]types.Task{
				prototypeTasks[5],
			},
		},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := getFilteredTasks(tt.flags)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i := range got {
				got[i].Deadline = got[i].Deadline.Round(time.Second)
				tt.want[i].Deadline = tt.want[i].Deadline.Round(time.Second)

				if !reflect.DeepEqual(tt.want[i], got[i]) {
					t.Errorf("got %v, want %v", got, tt.want)
				}
			}
		})
	}

	err := os.Remove("./bril.json")
	if err != nil {
		t.Error(err)
	}
}
