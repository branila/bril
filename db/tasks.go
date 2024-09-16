package db

import (
	"github.com/branila/bril/types"
)

func GetTasks() []types.Task {
	return db.Tasks
}

func CreateTask(t types.Task) {
	db.Tasks = append(db.Tasks, t)

	syncDb()
}
