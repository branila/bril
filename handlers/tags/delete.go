package tags

import (
	"fmt"
	"os"

	"github.com/branila/bril/db"
)

func Delete() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: bril tag delete <tag name>")
		return
	}

	tag := os.Args[3]

	tags := db.GetTags()

	for _, t := range tags {
		if t.Name == tag {
			db.DeleteTag(tag)
			fmt.Printf("Tag %s deleted\n", tag)
			return
		}
	}
}
