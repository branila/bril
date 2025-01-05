package finder

import (
	"fmt"
	"os"
	"strings"

	lister "github.com/branila/bril/handlers/list"
	"github.com/branila/bril/types"
)

func Find() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril find <query> [--all | -a] [--done | -d] [--tag <string> | -t <string>] [--expired | -e] [--deleted | -x] [--sort <name|priority|deadline> | -s <name|priority|deadline>]")
		return
	}

	flags := lister.GetFlags(os.Args[3:])

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
	query = strings.ToLower(query)

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
