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
		fmt.Println("Usage: bril find <query>")
		return
	}

	query := os.Args[2]

	tasks, err := searchTasks(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(tasks) == 0 {
		fmt.Printf("No tasks found for query %q\n", query)
		return
	}

	printTasks(tasks, query)
}

func searchTasks(query string) ([]types.Task, error) {
	query = strings.ToLower(os.Args[2])

	flags := lister.GetFlags(os.Args[3:])
	tasks := lister.GetFilteredTasks(flags)

	sortedTasks, err := lister.SortTsks(tasks, flags.Sort)
	if err != nil {
		return nil, err
	}

	foundTasks := []types.Task{}

	for _, task := range sortedTasks {
		taskName := strings.ToLower(task.Name)

		if strings.Contains(taskName, query) {
			foundTasks = append(foundTasks, task)
		}
	}

	return foundTasks, nil
}

func printTasks(tasks []types.Task, query string) {
	fmt.Printf("Tasks found for query %q:\n\n", query)

	for _, task := range tasks {
		lister.PrintTaskBrief(task)
	}
}
