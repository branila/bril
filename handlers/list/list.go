package lister

import (
	"fmt"
)

func List() {
	flags := getFlags()

	filteredTasks := getFilteredTasks(flags)

	if len(filteredTasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	sortedTasks, err := sortTasks(filteredTasks, flags.Sort)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Task list:\n\n")

	for _, task := range sortedTasks {
		PrintTaskBrief(task)
	}
}
