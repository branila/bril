package types

import (
	"time"
)

type Tag struct {
	Name     string
	Priority int // Sefault priority for tasks with this tag
}

type Note struct {
	Text     string
	Creation time.Time
}

type Task struct {
	Id         int
	Name       string
	Done       bool
	Priority   int
	Tag        Tag
	Notes      []Note
	Creation   time.Time
	Completion time.Time
	Expiration time.Time
	Elapsed    int // Seconds spent on the task
}

type Db struct {
	tasks []Task
	tags  []Tag // List of user-defined tags
}
