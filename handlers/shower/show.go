package shower

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/branila/bril/db"
)

func Show() {
	tasks := db.GetTasks()

	var taskId int

	if len(os.Args) < 3 {
		fmt.Println("Usage: bril show <task_id>")
		return
	}

	taskId, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Invalid task id")
	}

	for _, task := range tasks {
		if task.Id == taskId && !task.Deleted {
			printTask(task)
			return
		}
	}

	log.Fatal("Task not found")
}
