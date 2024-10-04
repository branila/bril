package adder

import "log"

func getPriority(setPriority, tagPriority int) int {
	isPrioritySet := setPriority != -1
	hasTagPriority := tagPriority != -1

	if !isPrioritySet && !hasTagPriority {
		return 0
	}

	if setPriority < 0 || setPriority > 10 {
		log.Fatal("Priority must be in the range of 0-10")
	}

	priority := setPriority

	if !isPrioritySet && hasTagPriority {
		priority = tagPriority
	}

	return priority
}
