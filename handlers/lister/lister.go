package lister

import (
	"fmt"

	"github.com/branila/bril/db"
)

func List() {
	tasks := db.GetTasks()

	fmt.Print("\nTask list:\n\n")

	for _, task := range tasks {
		if task.Done || task.Deleted {
			continue
		}

		fmt.Printf("[Id: %d] %s\n", task.Id, task.Name)
	}

	fmt.Println()
}
