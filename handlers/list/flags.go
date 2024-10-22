package handlers

import "flag"

type ListFlags struct {
	All  bool
	Done bool
}

func getFlags() ListFlags {
	var flags ListFlags

	flag.BoolVar(&flags.All, "a", false, "Show all tasks")
	flag.BoolVar(&flags.All, "all", false, "Show all tasks")

	flag.BoolVar(&flags.Done, "d", false, "Show done tasks")
	flag.BoolVar(&flags.Done, "done", false, "Show done tasks")

	flag.Parse()

	return flags
}
