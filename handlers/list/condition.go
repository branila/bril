package lister

import (
	"reflect"
	"strings"
	"time"

	"github.com/branila/bril/types"
)

func shouldPrint(task types.Task, flags ListFlags) bool {
	if reflect.DeepEqual(flags, ListFlags{}) {
		return !task.Done && !task.Deleted
	}

	showAll := flags.All
	showDone := flags.Done && task.Done
	showDeleted := flags.Deleted && task.Deleted
	showExpired := flags.Expired && time.Now().After(task.Deadline)

	tagMatch := strings.EqualFold(task.Tag, flags.Tag) && !task.Done && !task.Deleted

	return showAll || showDone || showDeleted || showExpired || tagMatch
}
