package shower

import (
	"fmt"

	"github.com/branila/bril/types"
)

func printTask(task types.Task) {
	fmt.Printf("\n[%d] %s\n", task.Id, task.Name)

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
