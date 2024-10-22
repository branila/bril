package undoer

import (
	"fmt"
	"os"
	"strconv"

	"github.com/branila/bril/db"
)

func Undo() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril undo <task id>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("'%s' is not a valid task id\n", os.Args[2])
		return
	}

	if id < 0 || id >= len(db.GetTasks()) {
		fmt.Printf("Task %d not found", id)
		return
	}

	tasks := db.GetTasks()

	if !tasks[id].Done {
		fmt.Printf("Task %d already undone\n", id)
		return
	}

	tasks[id].Done = false
	tasks[id].Deleted = false

	db.SyncDb()

	fmt.Printf("Task %d marked as undone\n", id)
}
