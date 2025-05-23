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
		fmt.Println("Usage: bril add <task name> [--priority <int> | -p <int>] [--due <date> | -d <date>] [--tag <string> | -t <string>] [--note <string> | -n <string>]")
		return
	}

	flags := getFlags()

	tag := getTag(flags.Tag)

	if flags.Tag != "" && tag.Name == "" {
		fmt.Printf("Tag '%s' not found\n", flags.Tag)
		return
	}

	deadline, err := getDeadline(flags.Due)
	if err != nil {
		fmt.Println(err)
		return
	}

	priority, err := getPriority(flags.Priority, tag.Priority)
	if err != nil {
		fmt.Println(err)
		return
	}

	task := types.Task{
		Id:       len(db.GetTasks()),
		Name:     os.Args[2],
		Priority: priority,
		Tag:      tag.Name,
		Note:     flags.Note,
		Creation: time.Now(),
		Deadline: deadline,
	}

	db.CreateTask(task)

	fmt.Printf("Created task '%s' (Id: %d)\n", task.Name, task.Id)
}
