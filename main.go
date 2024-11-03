package main

import (
	"fmt"
	"os"

	"github.com/branila/bril/db"
	adder "github.com/branila/bril/handlers/add"
	deleter "github.com/branila/bril/handlers/delete"
	doer "github.com/branila/bril/handlers/done"
	lister "github.com/branila/bril/handlers/list"
	reminder "github.com/branila/bril/handlers/remind"
	restorer "github.com/branila/bril/handlers/restore"
	undoer "github.com/branila/bril/handlers/undo"
	viewer "github.com/branila/bril/handlers/view"
)

func main() {
	db.Init()

	if len(os.Args) < 2 {
		fmt.Println("Usage: bril <command> [args]")
		return
	}

	switch os.Args[1] {
	case "add":
		adder.Add()

	case "list":
		lister.List()

	case "view":
		viewer.View()

	case "done":
		doer.Done()

	case "undo":
		undoer.Undo()

	case "delete":
		deleter.Delete()

	case "restore":
		restorer.Delete()

	case "remind":
		reminder.Remind()

	default:
		fmt.Println("Unknown command")
	}
}
