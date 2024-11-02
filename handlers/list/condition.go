package lister

import (
	"reflect"
	"strings"
	"time"

	"github.com/branila/bril/types"
)

func shouldPrint(task types.Task, flags ListFlags) bool {
	// If no specific flags are set, only show tasks that are not done or deleted
	if reflect.DeepEqual(flags, ListFlags{}) {
		return !task.Done && !task.Deleted
	}

	// Check for the Done flag
	if flags.Done && !task.Done {
		return false
	}
	// Check for the Deleted flag
	if flags.Deleted && !task.Deleted {
		return false
	}
	// Check for the Expired flag
	if flags.Expired && (task.Deadline.IsZero() || !task.Deadline.Before(time.Now())) {
		return false
	}
	// Check for Tag filtering
	if flags.Tag != "" && !strings.EqualFold(task.Tag, flags.Tag) {
		return false
	}
	// If the All flag is set, return true for all tasks
	if flags.All {
		return true
	}

	// If none of the above conditions are false, the task should be printed
	return true
}
