package handlers

import (
	"fmt"
	"os"

	"github.com/branila/bril/db"
)

func List() {
	tasks := db.GetTasks()

	// Remove non-flag arguments
	os.Args = os.Args[1:]

	flags := getFlags()

	fmt.Print("\nTask list:\n\n")

	for _, task := range tasks {

		if flags.Done && !task.Done {
			continue
		}

		if !flags.Done && task.Done && !flags.All {
			continue
		}

		fmt.Printf("[Id: %d] %s", task.Id, task.Name)

		if task.Done {
			fmt.Print(" (done")

			if task.Deleted {
				fmt.Print(", deleted")
			}

			fmt.Print(")")

		} else if task.Deleted {
			fmt.Print(" (deleted)")
		}

		fmt.Println()
	}

	fmt.Println()
}
