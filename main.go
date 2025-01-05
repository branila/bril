package main

import (
	"fmt"
	"os"

	"github.com/branila/bril/db"
	adder "github.com/branila/bril/handlers/add"
	deleter "github.com/branila/bril/handlers/delete"
	doer "github.com/branila/bril/handlers/done"
	modifier "github.com/branila/bril/handlers/edit"
	finder "github.com/branila/bril/handlers/find"
	lister "github.com/branila/bril/handlers/list"
	reminder "github.com/branila/bril/handlers/remind"
	resetter "github.com/branila/bril/handlers/reset"
	restorer "github.com/branila/bril/handlers/restore"
	tags "github.com/branila/bril/handlers/tags"
	undoer "github.com/branila/bril/handlers/undo"
	viewer "github.com/branila/bril/handlers/view"
)

func main() {
	db.Init()

	fmt.Println()
	defer fmt.Println()

	if len(os.Args) < 2 {
		fmt.Println("Usage: bril <command> [args]\n\nDocs: https://github.com/branila/bril/blob/main/README.md")
		return
	}

	switch os.Args[1] {
	case "add":
		adder.Add()

	case "ls", "list":
		lister.List()

	case "view":
		viewer.View()

	case "do", "done":
		doer.Done()

	case "undo":
		undoer.Undo()

	case "rm", "delete":
		deleter.Delete()

	case "restore":
		restorer.Delete()

	case "remind":
		reminder.Remind()

	case "find":
		finder.Find()

	case "reset":
		resetter.Reset()

	case "edit":
		modifier.Edit()

	case "tags", "tag":
		tags.Handle()

	default:
		fmt.Println("Unknown command")
	}
}
