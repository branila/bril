package main

import (
	"fmt"
	"os"

	"github.com/branila/bril/db"
	adder "github.com/branila/bril/handlers/add"
	deleter "github.com/branila/bril/handlers/delete"
	doer "github.com/branila/bril/handlers/done"
	finder "github.com/branila/bril/handlers/find"
	lister "github.com/branila/bril/handlers/list"
	reminder "github.com/branila/bril/handlers/remind"
	resetter "github.com/branila/bril/handlers/reset"
	restorer "github.com/branila/bril/handlers/restore"
	undoer "github.com/branila/bril/handlers/undo"
	viewer "github.com/branila/bril/handlers/view"
)

func main() {
	db.Init()

	fmt.Println()
	defer fmt.Println()

	if len(os.Args) < 2 {
		fmt.Println("Usage: bril <command> [args]")
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

	default:
		fmt.Println("Unknown command")
	}
}
