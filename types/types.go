package types

import (
	"time"
)

type Tag struct {
	Name     string
	Priority int
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
	Elapsed    int
}
