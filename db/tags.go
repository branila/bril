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

func DeleteTag(name string) {
	for i, t := range db.Tags {
		if t.Name == name {
			db.Tags = append(db.Tags[:i], db.Tags[i+1:]...)

			if err := SyncDb(); err != nil {
				fmt.Printf("Error deleting tag (%q)\n", err)
			}
		}
	}

	for i, t := range db.Tasks {
		if t.Tag == name {
			db.Tasks[i].Tag = ""
		}
	}
}

func UpdateTag(t types.Tag) {
	for i, tag := range db.Tags {
		if tag.Name == t.Name {
			db.Tags[i] = t

			if err := SyncDb(); err != nil {
				fmt.Printf("Error updating tag (%q)\n", err)
			}
		}
	}
}
