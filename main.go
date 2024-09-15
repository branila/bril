package main

import (
	"fmt"
	"time"

	"github.com/branila/bril/types"
)

func main() {
	t := types.Task{
		Id:       0,
		Name:     "Create an advanced terminal-centric task manager",
		Done:     false,
		Priority: 100,
		Tag:      types.Tag{},
		Notes: []types.Note{
			{
				Text:     "This is a note",
				Creation: time.Now(),
			},
		},
		Creation:   time.Now(),
		Completion: time.Time{},
		Expiration: time.Time{},
		Elapsed:    0,
	}

	fmt.Println(t)
}
