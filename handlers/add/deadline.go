package adder

import (
	"strings"
	"time"

	utils "github.com/branila/bril/utils/parse"
)

func getDeadline(due string) (time.Time, error) {
	if due == "" {
		return time.Time{}, nil
	}
	switch strings.ToLower(due) {
	case "today":
		return time.Now().Truncate(24 * time.Hour), nil
	case "tomorrow":
		return time.Now().Truncate(24 * time.Hour).Add(24 * time.Hour), nil
	}

	if utils.IsDurationFormat(due) {
		return utils.ParseDuration(due)
	}

	return utils.ParseDate(due)
}
