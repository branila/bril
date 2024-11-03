package adder

import (
	"time"

	utils "github.com/branila/bril/utils/parse"
)

func getDeadline(due string) (time.Time, error) {
	if due == "" {
		return time.Time{}, nil
	}

	if utils.IsDurationFormat(due) {
		return utils.ParseDuration(due)
	}

	return utils.ParseDate(due)
}
