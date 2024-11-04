package lister

import (
	"fmt"
	"os"
)

func List() {
	flags := GetFlags(os.Args[2:])

	filteredTasks := GetFilteredTasks(flags)

	if len(filteredTasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	sortedTasks, err := SortTsks(filteredTasks, flags.Sort)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Task list:\n\n")

	for _, task := range sortedTasks {
		PrintTaskBrief(task)
	}
}
