package lister

import (
	"fmt"
	"os"
)

func List() {
	// Remove non-flag arguments
	os.Args = os.Args[1:]

	flags := getFlags()

	filteredTasks := getFilteredTasks(flags)

	if len(filteredTasks) == 0 {
		fmt.Println("No tasks found matching the criteria")
		return
	}

	fmt.Print("Task list:\n\n")

	for _, task := range filteredTasks {
		PrintTaskBrief(task)
	}
}
