package adder

import (
	"fmt"
	"os"
	"time"

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

	tag := getTag(flags.Tag)

	if flags.Tag != "" && tag.Name == "" {
		fmt.Printf("Tag '%s' not found\n", flags.Tag)
		return
	}

	task := types.Task{
		Id:       len(db.GetTasks()),
		Name:     name,
		Priority: getPriority(flags.Priority, tag.Priority),
		Tag:      tag.Name,
		Note:     flags.Note,
		Creation: time.Now(),
		Deadline: getDeadline(flags.Due),
	}

	db.CreateTask(task)

	fmt.Printf("Created task '%s' (Id: %d)\n", task.Name, task.Id)
}
