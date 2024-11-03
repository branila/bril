package lister

import (
	"strings"
	"time"

	"github.com/branila/bril/types"
)

func noFilteringFlagSet(flags ListFlags) bool {
	return !flags.All && !flags.Done && !flags.Deleted && !flags.Expired && flags.Tag == ""
}

func shouldPrint(task types.Task, flags ListFlags) bool {
	if noFilteringFlagSet(flags) {
		return !task.Done && !task.Deleted
	}

	if flags.Done && !task.Done {
		return false
	}

	if flags.Deleted && !task.Deleted {
		return false
	}

	if flags.Expired && (task.Deadline.IsZero() || !task.Deadline.Before(time.Now())) {
		return false
	}

	if flags.Tag != "" && !strings.EqualFold(task.Tag, flags.Tag) {
		return false
	}

	if flags.All {
		return true
	}

	return true
}
