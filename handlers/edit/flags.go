package modifier

import (
	"flag"
	"os"

	"github.com/branila/bril/utils"
)

type EditFlags struct {
	Name     string
	State    string
	Tag      string
	Priority int
	Note     string
	Deadline string
}

func (e EditFlags) String() string {
	return utils.PrettifyObject(e)
}

func getFlags() EditFlags {
	fs := flag.NewFlagSet("editFlags", flag.ExitOnError)

	var flags EditFlags

	fs.StringVar(&flags.Name, "name", "", "Set the name of the task")
	fs.StringVar(&flags.Name, "n", "", "Set the name of the task")

	fs.StringVar(&flags.State, "state", "", "Set the state of the task")
	fs.StringVar(&flags.State, "s", "", "Set the tag of the task")

	fs.StringVar(&flags.Tag, "tag", "", "Set the tag of the task")
	fs.StringVar(&flags.Tag, "t", "", "Set the priority of the task")

	fs.IntVar(&flags.Priority, "priority", -1, "Set the priority of the task")
	fs.IntVar(&flags.Priority, "p", -1, "Set the note of the task")

	fs.StringVar(&flags.Note, "note", "", "Set the note of the task")

	fs.StringVar(&flags.Deadline, "deadline", "", "Set the deadline of the task")
	fs.StringVar(&flags.Deadline, "d", "", "Set the deadline of the task")

	fs.Parse(os.Args[3:])

	return flags
}
