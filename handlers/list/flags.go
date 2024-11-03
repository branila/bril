package lister

import (
	"flag"
	"os"

	"github.com/branila/bril/utils"
)

type ListFlags struct {
	All     bool
	Done    bool
	Tag     string
	Expired bool
	Deleted bool
	Sort    string
}

func (f ListFlags) String() string {
	return utils.PrettifyObject(f)
}

func getFlags() ListFlags {
	var flags ListFlags

	fs := flag.NewFlagSet("listFlags", flag.ExitOnError)

	fs.BoolVar(&flags.All, "a", false, "Show all tasks")
	fs.BoolVar(&flags.All, "all", false, "Show all tasks")

	fs.BoolVar(&flags.Done, "d", false, "Show done tasks")
	fs.BoolVar(&flags.Done, "done", false, "Show done tasks")

	fs.StringVar(&flags.Tag, "t", "", "Show tasks with a specific tag")
	fs.StringVar(&flags.Tag, "tag", "", "Show tasks with a specific tag")

	fs.BoolVar(&flags.Expired, "e", false, "Only show expired tasks")
	fs.BoolVar(&flags.Expired, "expired", false, "Only show expired tasks")

	fs.BoolVar(&flags.Deleted, "del", false, "Show deleted tasks")
	fs.BoolVar(&flags.Deleted, "deleted", false, "Show deleted tasks")

	fs.StringVar(&flags.Sort, "s", "", "Sort tasks by a specific field")
	fs.StringVar(&flags.Sort, "sort", "", "Sort tasks by a specific field")

	fs.Parse(os.Args[2:])

	return flags
}
