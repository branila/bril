package doer

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/branila/bril/db"
)

func Done() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril done <task id>")
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

	if tasks[id].Done {
		fmt.Printf("Task %d already done\n", id)
		return
	}

	tasks[id].Done = true
	tasks[id].Deleted = false
	tasks[id].Completion = time.Now()

	db.SyncDb()

	fmt.Printf("Task %d marked as done\n", id)
}
