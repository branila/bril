package db

import (
	"github.com/branila/bril/types"
)

func GetTags() []types.Tag {
	return db.Tags
}

func CreateTag(t types.Tag) {
	db.Tags = append(db.Tags, t)

	SyncDb()
}
