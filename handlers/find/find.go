package finder

import (
	"fmt"
	"os"
	"strings"

	lister "github.com/branila/bril/handlers/list"
	"github.com/branila/bril/types"
)

func Find() {
	flags := lister.GetFlags(os.Args[3:])

	if len(os.Args) < 3 {
		fmt.Println("Usage: bril find <query>")
		return
	}

	query := os.Args[2]

	tasks := lister.GetFilteredTasks(flags)

	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	tasks, err := lister.SortTasks(tasks, flags.Sort)
	if err != nil {
		fmt.Println(err)
		return
	}

	tasks = searchTasks(tasks, query)

	if len(tasks) == 0 {
		fmt.Printf("No tasks found for query %q\n", query)
		return
	}

	printTasks(tasks, query)
}

func searchTasks(tasks []types.Task, query string) []types.Task {
	query = strings.ToLower(os.Args[2])

	foundTasks := []types.Task{}

	for _, task := range tasks {
		taskName := strings.ToLower(task.Name)

		if strings.Contains(taskName, query) {
			foundTasks = append(foundTasks, task)
		}
	}

	return foundTasks
}

func printTasks(tasks []types.Task, query string) {
	fmt.Printf("Tasks found for query %q:\n\n", query)

	for _, task := range tasks {
		lister.PrintTaskBrief(task)
	}
}
