package db

import (
	"fmt"

	"github.com/branila/bril/types"
)

func GetTags() []types.Tag {
	return db.Tags
}

func CreateTag(t types.Tag) {
	db.Tags = append(db.Tags, t)

	if err := SyncDb(); err != nil {
		fmt.Printf("Error creating tag (%q)\n", err)
	}
}
