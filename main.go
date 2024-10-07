package main

import (
	"fmt"
	"os"

	"github.com/branila/bril/db"
	addHandler "github.com/branila/bril/handlers/add"
	doneHandler "github.com/branila/bril/handlers/done"
	listHandler "github.com/branila/bril/handlers/list"
	viewHandler "github.com/branila/bril/handlers/view"
)

func main() {
	db.Init()

	if len(os.Args) < 2 {
		fmt.Println("Usage: bril <command> [args]")
		return
	}

	switch os.Args[1] {
	case "add":
		addHandler.Add()

	case "list":
		listHandler.List()

	case "view":
		viewHandler.View()

	case "done":
		doneHandler.Done()

	default:
		fmt.Println("Unknown command")
	}
}
