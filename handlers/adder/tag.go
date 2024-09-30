package adder

import (
	"log"
	"strings"

	"github.com/branila/bril/db"
	"github.com/branila/bril/types"
)

func getTag(tag string) types.Tag {
	if tag == "" {
		return types.Tag{
			Name:     "",
			Priority: -1,
		}
	}

	tags := db.GetTags()

	for _, t := range tags {
		if strings.ToLower(tag) == t.Name {
			return t
		}
	}

	log.Fatalf("Tag %s not found", tag)

	return types.Tag{}
}
