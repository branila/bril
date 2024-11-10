package tags

import (
	"flag"
	"fmt"
	"os"

	"github.com/branila/bril/db"
)

type EditTagFlags struct {
	Name     string
	Priority int
}

func getEditFlags() EditTagFlags {
	fs := flag.NewFlagSet("editTag", flag.ExitOnError)

	var flags EditTagFlags

	fs.StringVar(&flags.Name, "name", "", "Name of the tag")
	fs.StringVar(&flags.Name, "n", "", "Name of the tag")
	fs.IntVar(&flags.Priority, "priority", -1, "Priority of the tag")
	fs.IntVar(&flags.Priority, "p", -1, "Priority of the tag")

	fs.Parse(os.Args[4:])

	return flags
}

func Edit() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: bril tag edit <tag>")
		return
	}

	tag := os.Args[3]

	flags := getEditFlags()

	tags := db.GetTags()

	for i, t := range tags {
		if t.Name == tag {
			if flags.Name != "" {
				tags[i].Name = flags.Name
			}

			if flags.Priority >= 0 {
				tags[i].Priority = flags.Priority
			}

			db.UpdateTag(tags[i])

			fmt.Printf("Tag %s edited\n", tag)

			return
		}
	}
}
