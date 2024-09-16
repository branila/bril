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
	TagId      int           `json:"tagId"`
	Notes      []Note        `json:"notes"`
	Creation   time.Time     `json:"creation"`
	Completion time.Time     `json:"completion"`
	Expiration time.Time     `json:"expiration"`
	Timer      time.Duration `json:"timer"`
}

func (t Task) String() string {
	return utils.PrettifyObject(t)
}

type Note struct {
	Text     string    `json:"text"`
	Creation time.Time `json:"creation"`
}

func (n Note) String() string {
	return utils.PrettifyObject(n)
}
