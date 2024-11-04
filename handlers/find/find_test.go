package finder

import (
	"testing"

	"github.com/branila/bril/types"
)

func TestSearchTasks(t *testing.T) {
	tasks := []types.Task{
		{Id: 0, Name: "Go to the gym"},
		{Id: 1, Name: "Go to the groceries store"},
		{Id: 2, Name: "Go to the house"},
		{Id: 3, Name: "Clean the house"},
		{Id: 4, Name: "Clean the gym"},
		{Id: 5, Name: "Clean the groceries store"},
		{Id: 6, Name: "Buy groceries"},
		{Id: 7, Name: "Buy the gym"},
		{Id: 8, Name: "Buy the house"},
	}

	tests := []struct {
		query string
		want  []types.Task
	}{
		{
			query: "go",
			want: []types.Task{
				{Id: 0, Name: "Go to the gym"},
				{Id: 1, Name: "Go to the groceries store"},
				{Id: 2, Name: "Go to the house"},
			},
		},
		{
			query: "clean",
			want: []types.Task{
				{Id: 3, Name: "Clean the house"},
				{Id: 4, Name: "Clean the gym"},
				{Id: 5, Name: "Clean the groceries store"},
			},
		},
		{
			query: "buy",
			want: []types.Task{
				{Id: 6, Name: "Buy groceries"},
				{Id: 7, Name: "Buy the gym"},
				{Id: 8, Name: "Buy the house"},
			},
		},
		{
			query: "house",
			want: []types.Task{
				{Id: 2, Name: "Go to the house"},
				{Id: 3, Name: "Clean the house"},
				{Id: 8, Name: "Buy the house"},
			},
		},
		{
			query: "gym",
			want: []types.Task{
				{Id: 0, Name: "Go to the gym"},
				{Id: 4, Name: "Clean the gym"},
				{Id: 7, Name: "Buy the gym"},
			},
		},
		{
			query: "groceries",
			want: []types.Task{
				{Id: 1, Name: "Go to the groceries store"},
				{Id: 5, Name: "Clean the groceries store"},
				{Id: 6, Name: "Buy groceries"},
			},
		},
		{
			query: "store",
			want: []types.Task{
				{Id: 1, Name: "Go to the groceries store"},
				{Id: 5, Name: "Clean the groceries store"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			got := searchTasks(tasks, tt.query)

			if len(got) != len(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("Got %v, want %v", got, tt.want)
				}
			}
		})
	}
}
