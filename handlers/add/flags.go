package adder

import (
	"flag"
	"os"

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

	fs := flag.NewFlagSet("addFlags", flag.ExitOnError)

	fs.IntVar(&flags.Priority, "p", -1, "Task Priority")
	fs.IntVar(&flags.Priority, "priority", -1, "Task Priority")

	fs.StringVar(&flags.Due, "d", "", "Task Due date")
	fs.StringVar(&flags.Due, "due", "", "Task Due date")
	fs.StringVar(&flags.Due, "deadline", "", "Task Due date")

	fs.StringVar(&flags.Tag, "t", "", "Task Tag")
	fs.StringVar(&flags.Tag, "tag", "", "Task Tag")

	fs.StringVar(&flags.Note, "n", "", "Task Note")
	fs.StringVar(&flags.Note, "note", "", "Task Note")

	fs.Parse(os.Args[3:])

	return flags
}
