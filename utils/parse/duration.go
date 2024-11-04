package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func IsDurationFormat(inputDeadline string) bool {
	durationRegex := regexp.MustCompile(`^(?:\d+[yMwdhms])+$`)

	return durationRegex.MatchString(inputDeadline)
}

func ParseDuration(inputDeadline string) (time.Time, error) {
	now := time.Now()
	durationMap := map[string]int{
		"y": 0, "M": 0, "w": 0, "d": 0,
		"h": 0, "m": 0, "s": 0,
	}

	re := regexp.MustCompile(`(\d+)([yMwdhms])`)
	matches := re.FindAllStringSubmatch(inputDeadline, -1)

	if len(matches) == 0 {
		return time.Time{}, fmt.Errorf("no valid duration found")
	}

	for _, match := range matches {
		value, err := strconv.Atoi(match[1])
		if err != nil {
			return time.Time{}, fmt.Errorf("invalid number: %s", match[1])
		}
		unit := match[2]
		durationMap[unit] += value
	}

	deadline := now.AddDate(
		durationMap["y"],
		durationMap["M"],
		durationMap["w"]*7+durationMap["d"],
	)

	deadline = deadline.Add(
		time.Hour*time.Duration(durationMap["h"]) +
			time.Minute*time.Duration(durationMap["m"]) +
			time.Second*time.Duration(durationMap["s"]),
	)

	return deadline, nil
}
