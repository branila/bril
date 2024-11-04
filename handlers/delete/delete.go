package deleter

import (
	"fmt"
	"os"
	"strconv"

	"github.com/branila/bril/db"
)

func Delete() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril delete <task id>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil || id < 0 {
		fmt.Printf("'%s' is not a valid task id\n", os.Args[2])
		return
	}

	if id >= len(db.GetTasks()) {
		fmt.Printf("Task %d not found", id)
		return
	}

	tasks := db.GetTasks()

	if tasks[id].Deleted {
		fmt.Printf("Task %d has already been deleted\n", id)
		return
	}

	tasks[id].Deleted = true

	if err := db.SyncDb(); err != nil {
		fmt.Printf("Error deleting task (%q)\n", err)
		return
	}

	fmt.Printf("Task %d marked as deleted\n", id)
}
