package lister

import (
	"fmt"
	"sort"

	"github.com/branila/bril/types"
)

func SortTasks(tasks []types.Task, sortFlag string) ([]types.Task, error) {
	if sortFlag == "" {
		return tasks, nil
	}

	switch sortFlag {
	case "name":
		return orderByName(tasks), nil
	case "deadline":
		return orderByDeadline(tasks), nil
	case "priority":
		return orderByPriority(tasks), nil
	default:
		return tasks, fmt.Errorf("invalid order flag: %s", sortFlag)
	}
}

func orderByName(tasks []types.Task) []types.Task {
	orderedTasks := make([]types.Task, len(tasks))
	copy(orderedTasks, tasks)

	sort.Slice(orderedTasks, func(i, j int) bool {
		return orderedTasks[i].Name < orderedTasks[j].Name
	})

	return orderedTasks
}

func orderByDeadline(tasks []types.Task) []types.Task {
	orderedTasks := make([]types.Task, len(tasks))
	copy(orderedTasks, tasks)

	sort.Slice(orderedTasks, func(i, j int) bool {
		return orderedTasks[i].Deadline.Before(orderedTasks[j].Deadline)
	})

	return orderedTasks
}

func orderByPriority(tasks []types.Task) []types.Task {
	orderedTasks := make([]types.Task, len(tasks))
	copy(orderedTasks, tasks)

	sort.Slice(orderedTasks, func(i, j int) bool {
		return orderedTasks[i].Priority > orderedTasks[j].Priority
	})

	return orderedTasks
}
