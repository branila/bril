package viewer

import (
	"fmt"

	"github.com/branila/bril/types"
)

func printTask(task types.Task) {
	fmt.Printf("\n[%d] %s", task.Id, task.Name)

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

	if task.Note != "" {
		fmt.Printf("  - Note: %s\n", task.Note)
	}

	fmt.Println()
}
