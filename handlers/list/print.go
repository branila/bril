package lister

import (
	"fmt"
	"strings"
	"time"

	"github.com/branila/bril/types"
)

func PrintTaskBrief(task types.Task) {
	fmt.Printf("[Id: %d] %s", task.Id, task.Name)

	var states []string

	if task.Done {
		states = append(states, "done")
	}
	if task.Deleted {
		states = append(states, "deleted")
	}
	if time.Now().After(task.Deadline) && !task.Deadline.IsZero() {
		states = append(states, "expired")
	}

	if len(states) > 0 {
		fmt.Printf(" (%s)", strings.Join(states, ", "))
	}

	fmt.Println()
}
