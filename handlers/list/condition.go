package lister

import (
	"strings"
	"time"

	"github.com/branila/bril/types"
)

func shouldPrint(task types.Task, flags ListFlags) bool {
	// Default behavior
	if !flags.All && !flags.Done && !flags.Deleted && !flags.Expired && flags.Tag == "" {
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

	return true
}
