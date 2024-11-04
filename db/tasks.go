package db

import (
	"fmt"

	"github.com/branila/bril/types"
)

func GetTasks() []types.Task {
	return db.Tasks
}

func GetTask(id int) (types.Task, bool) {
	for _, t := range db.Tasks {
		if t.Id == id {
			return t, true
		}
	}

	return types.Task{}, false
}

func CreateTask(t types.Task) {
	db.Tasks = append(db.Tasks, t)

	if err := SyncDb(); err != nil {
		fmt.Printf("Error creating task (%q)\n", err)
	}
}

func SetTasks(tasks []types.Task) {
	db.Tasks = tasks

	if err := SyncDb(); err != nil {
		fmt.Printf("Error setting tasks (%q)\n", err)
	}
}
