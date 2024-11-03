package reminder

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/branila/bril/db"
	utils "github.com/branila/bril/utils/parse"
)

func Remind() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril remind <task id> <time>")
		return
	}

	id, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid task id")
		return
	}

	deadline, err := parseDeadline(os.Args[3])
	if err != nil {
		fmt.Println("Invalid time format")
		return
	}

	timeToDeadline := deadline.Sub(time.Now())
	if timeToDeadline <= 0 {
		fmt.Println("The specified time is in the past")
		return
	}

	task, found := db.GetTask(id)
	if !found {
		fmt.Println("Task not found")
		return
	}

	err = startReminder(timeToDeadline, id, task.Name)
	if err != nil {
		fmt.Println("Failed to start reminder")
		return
	}

	fmt.Printf("Reminding task %d in %s\n", id, formatDuration(timeToDeadline))
}

func parseDeadline(timeInput string) (time.Time, error) {
	if utils.IsDurationFormat(timeInput) {
		return utils.ParseDuration(timeInput)
	}
	return utils.ParseDate(timeInput)
}

func startReminder(duration time.Duration, id int, taskName string) error {
	seconds := int(duration.Seconds())
	notifyCmd := fmt.Sprintf("sleep %d && notify-send 'Task %d' '%s'", seconds, id, taskName)
	cmd := exec.Command("bash", "-c", notifyCmd)
	return cmd.Start()
}

func formatDuration(d time.Duration) string {
	units := []struct {
		value int
		label string
	}{
		{int(d.Hours()) / 24, "days"},
		{int(d.Hours()) % 24, "hours"},
		{int(d.Minutes()) % 60, "minutes"},
		{int(d.Seconds()) % 60, "seconds"},
	}

	var parts []string

	for _, u := range units {
		if u.value > 0 {
			parts = append(parts, fmt.Sprintf("%d %s", u.value, u.label))
		}
	}

	if len(parts) == 0 {
		return "0 seconds"
	}

	if len(parts) == 1 {
		return parts[0]
	}

	return strings.Join(parts[:len(parts)-1], ", ") + " and " + parts[len(parts)-1]
}
