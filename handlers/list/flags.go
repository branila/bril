package lister

import (
	"flag"
)

type ListFlags struct {
	All     bool
	Done    bool
	Tag     string
	Expired bool
	Deleted bool
}

func getFlags() ListFlags {
	var flags ListFlags

	flag.BoolVar(&flags.All, "a", false, "Show all tasks")
	flag.BoolVar(&flags.All, "all", false, "Show all tasks")

	flag.BoolVar(&flags.Done, "d", false, "Show done tasks")
	flag.BoolVar(&flags.Done, "done", false, "Show done tasks")

	flag.StringVar(&flags.Tag, "t", "", "Show tasks with a specific tag")
	flag.StringVar(&flags.Tag, "tag", "", "Show tasks with a specific tag")

	flag.BoolVar(&flags.Expired, "e", false, "Only show expired tasks")
	flag.BoolVar(&flags.Expired, "expired", false, "Only show expired tasks")

	flag.BoolVar(&flags.Deleted, "del", false, "Show deleted tasks")
	flag.BoolVar(&flags.Deleted, "deleted", false, "Show deleted tasks")

	flag.Parse()

	return flags
}
