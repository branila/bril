package adder

import "fmt"

func getPriority(setPriority, tagPriority int) (int, error) {
	isPrioritySet := setPriority != -1
	hasTagPriority := tagPriority != -1

	if !isPrioritySet && !hasTagPriority {
		return 0
	}

	if setPriority < 0 || setPriority > 10 {
		return 0, fmt.Errorf("Priority must be in the range of 0-10")
	}

	priority := setPriority

	if !isPrioritySet && hasTagPriority {
		priority = tagPriority
	}

	return priority
}
