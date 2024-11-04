package adder

import (
	"time"

	utils "github.com/branila/bril/utils/parse"
)

func getDeadline(deadline string) (time.Time, error) {
	if deadline == "" {
		return time.Time{}, nil
	}

	if utils.IsRelativeFormat(deadline) {
		return utils.ParseRelativeFormat(deadline)
	}

	if utils.IsDurationFormat(deadline) {
		return utils.ParseDuration(deadline)
	}

	return utils.ParseDate(deadline)
}
