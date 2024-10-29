package adder

import (
	"strings"

	"github.com/branila/bril/db"
	"github.com/branila/bril/types"
)

func getTag(tag string) types.Tag {
	defaultTag := types.Tag{
		Name:     "",
		Priority: -1,
	}

	if tag == "" {
		return defaultTag
	}

	tags := db.GetTags()

	for _, t := range tags {
		if strings.ToLower(tag) == t.Name {
			return t
		}
	}

	return defaultTag
}
