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
	if err != nil {
		fmt.Printf("'%s' is not a valid task id\n", os.Args[2])
		return
	}

	if id < 0 || id >= len(db.GetTasks()) {
		fmt.Printf("Task %d not found", id)
		return
	}

	tasks := db.GetTasks()

	if tasks[id].Deleted {
		fmt.Printf("Task %d already deleted\n", id)
		return
	}

	tasks[id].Deleted = true

	db.SyncDb()

	fmt.Printf("Task %d marked as deleted\n", id)
}
