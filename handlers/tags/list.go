package tags

import (
	"fmt"

	"github.com/branila/bril/db"
	"github.com/branila/bril/types"
)

func PrintTagBrief(index int, tag types.Tag) {
	fmt.Printf("%d. %s", index+1, tag.Name)

	if tag.Priority >= 0 {
		fmt.Printf(" (Priority: %d)", tag.Priority)
	}

	fmt.Println()
}

func List() {
	tags := db.GetTags()

	if len(tags) == 0 {
		fmt.Println("No tags found")
		return
	}

	for i, tag := range tags {
		PrintTagBrief(i, tag)
	}
}
