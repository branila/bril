package adder

import (
	"time"

	"github.com/branila/bril/types"
)

func getNotes(note string) []types.Note {
	if note == "" {
		return nil
	}

	notes := []types.Note{
		{
			Text:     note,
			Creation: time.Now(),
		},
	}

	return notes
}
