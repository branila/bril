package types

import (
	"github.com/branila/bril/utils"
)

type Db struct {
	Tasks []Task
	Tags  []Tag
}

func (d Db) String() string {
	return utils.PrettifyObject(d)
}
