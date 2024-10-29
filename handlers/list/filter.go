package lister

import (
	"github.com/branila/bril/db"
	"github.com/branila/bril/types"
)

func getFilteredTasks(flags ListFlags) []types.Task {
	tasks := db.GetTasks()

	filteredTasks := []types.Task{}

	for _, task := range tasks {
		if shouldPrint(task, flags) {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return filteredTasks
}
