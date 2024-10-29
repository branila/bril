package adder

import (
	"flag"

	"github.com/branila/bril/utils"
)

type AddFlags struct {
	Priority int
	Due      string
	Tag      string
	Note     string
}

func (f AddFlags) String() string {
	return utils.PrettifyObject(f)
}

func getFlags() AddFlags {
	var flags AddFlags

	flag.IntVar(&flags.Priority, "p", -1, "Task Priority")
	flag.IntVar(&flags.Priority, "Priority", -1, "Task Priority")

	flag.StringVar(&flags.Due, "d", "", "Task Due date")
	flag.StringVar(&flags.Due, "Due", "", "Task Due date")
	flag.StringVar(&flags.Due, "deadline", "", "Task Due date")

	flag.StringVar(&flags.Tag, "t", "", "Task Tag")
	flag.StringVar(&flags.Tag, "Tag", "", "Task Tag")

	flag.StringVar(&flags.Note, "n", "", "Task Note")
	flag.StringVar(&flags.Note, "Note", "", "Task Note")

	flag.Parse()

	return flags
}
