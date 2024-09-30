package adder

import "log"

func getPriority(setPriority, tagPriority int) int {
	if setPriority < 0 || setPriority > 10 {
		log.Fatal("Priority must be in the range of 0-10")
	}

	priority := setPriority

	isPriorityUnset := priority == -1
	hasTagPriority := tagPriority != -1

	if isPriorityUnset && hasTagPriority {
		priority = tagPriority
	}

	return priority
}
