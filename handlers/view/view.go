package viewer

import (
	"fmt"
	"os"
	"strconv"

	"github.com/branila/bril/db"
)

func View() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril view <task id>")
		return
	}

	taskId, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid task id")
		return
	}

	task, found := db.GetTask(taskId)
	if !found {
		fmt.Println("Task not found")
		return
	}

	printTask(task)
}
