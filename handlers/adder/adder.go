package adder

import (
	"fmt"
	"os"

	"github.com/branila/bril/db"
	"github.com/branila/bril/types"
)

func Add() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril add <task name> [flags]")
		return
	}

	name := os.Args[2]

	// Remove non-flag arguments
	os.Args = os.Args[2:]

	flags := getFlags()

	tag := getTag(flags.tag)

	task := types.Task{
		Id:         len(db.GetTasks()),
		Name:       name,
		Priority:   getPriority(flags.priority, tag.Priority),
		Expiration: getExpiration(flags.due),
		Notes:      getNotes(flags.note),
		Tag:        tag.Name,
	}

	db.CreateTask(task)
}
