package types

import "github.com/branila/bril/utils"

type Tag struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"` // Default priority for tasks with this tag
}

func (t Tag) String() string {
	return utils.PrettifyObject(t)
}
