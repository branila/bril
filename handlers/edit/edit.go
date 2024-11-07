package modifier

import (
	"fmt"
	"os"
	"strconv"

	"github.com/branila/bril/db"
	"github.com/branila/bril/types"
	utils "github.com/branila/bril/utils/parse"
)

func allFlagsUnset(flags EditFlags) bool {
	return flags.Name == "" && flags.State == "" && flags.Tag == "" && flags.Priority == -1 && flags.Note == "" && flags.Deadline == ""
}

func Edit() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril edit <task id> [flags]")
		return
	}

	flags := getFlags()

	if allFlagsUnset(flags) {
		fmt.Println("Usage: bril edit <task id> [flags]")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid task id")
		return
	}

	tasks := db.GetTasks()

	if id < 0 || id >= len(tasks) {
		fmt.Println("Task not found")
		return
	}

	task, err := getEditedTask(tasks[id], flags)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks[id] = task

	db.SetTasks(tasks)

	fmt.Printf("Task %d edited successfully\n", id)
}

func getEditedTask(task types.Task, flags EditFlags) (types.Task, error) {
	if flags.Name != "" {
		task.Name = flags.Name
	}

	if flags.State != "" {
		switch flags.State {
		case "done":
			task.Done = true
		case "deleted":
			task.Deleted = true
		case "undone":
			task.Done = false
		case "active":
			task.Deleted = false
		default:
			return types.Task{}, fmt.Errorf("Invalid state")
		}
	}

	if flags.Tag != "" {
		task.Tag = flags.Tag
	}

	if flags.Priority != -1 {
		task.Priority = flags.Priority
	}

	if flags.Note != "" {
		task.Note = flags.Note
	}

	var err error

	if flags.Deadline != "" {
		if utils.IsDurationFormat(flags.Deadline) {
			task.Deadline, err = utils.ParseDuration(flags.Deadline)
		} else if utils.IsRelativeFormat(flags.Deadline) {
			task.Deadline, err = utils.ParseRelativeFormat(flags.Deadline)
		} else {
			task.Deadline, err = utils.ParseDate(flags.Deadline)
		}

		if err != nil {
			return types.Task{}, err
		}
	}

	return task, nil
}
