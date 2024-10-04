package lister

import (
	"fmt"

	"github.com/branila/bril/db"
)

func Show() {
	tasks := db.GetTasks()

	fmt.Print("\nLista delle task:\n\n")

	for _, task := range tasks {
		if task.Done || task.Deleted {
			continue
		}

		fmt.Printf("[%d] %s\n", task.Id, task.Name)

		if task.Priority != 0 {
			fmt.Printf("  - Priority: %d\n", task.Priority)
		}

		if task.Tag != "" {
			fmt.Printf("  - Tag: %s\n", task.Tag)
		}

		fmt.Printf("  - Creation: %s\n", task.Creation.Format("02-01-2006 15:04:05"))

		if !task.Deadline.IsZero() {
			fmt.Printf("  - Deadline: %s\n", task.Deadline.Format("02-01-2006 15:04:05"))
		}

		if len(task.Notes) > 0 {
			fmt.Println("  - Notes:")

			for _, note := range task.Notes {
				fmt.Printf("    - %s\n", note)
			}
		}

		fmt.Println()
	}
}
