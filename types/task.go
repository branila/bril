package types

import (
	"time"

	"github.com/branila/bril/utils"
)

type Task struct {
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	Done       bool          `json:"done"`
	Priority   int           `json:"priority"`
	Tag        string        `json:"tag"`
	Note       string        `json:"note"`
	Creation   time.Time     `json:"creation"`
	Completion time.Time     `json:"completion"`
	Deadline   time.Time     `json:"deadline"`
	Timer      time.Duration `json:"timer"`
	Deleted    bool          `json:"deleted"`
}

func (t Task) String() string {
	return utils.PrettifyObject(t)
}
