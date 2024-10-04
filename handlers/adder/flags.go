package adder

import (
	"flag"

	"github.com/branila/bril/utils"
)

type AdderFlags struct {
	priority int
	due      string
	tag      string
	note     string
}

func (f AdderFlags) String() string {
	return utils.PrettifyObject(f)
}

func getFlags() AdderFlags {
	var flags AdderFlags

	flag.IntVar(&flags.priority, "p", -1, "Task priority")
	flag.IntVar(&flags.priority, "priority", -1, "Task priority")

	flag.StringVar(&flags.due, "d", "", "Task due date")
	flag.StringVar(&flags.due, "due", "", "Task due date")
	flag.StringVar(&flags.due, "deadline", "", "Task due date")

	flag.StringVar(&flags.tag, "t", "", "Task tag")
	flag.StringVar(&flags.tag, "tag", "", "Task tag")

	flag.StringVar(&flags.note, "n", "", "Task note")
	flag.StringVar(&flags.note, "note", "", "Task note")

	flag.Parse()

	return flags
}
