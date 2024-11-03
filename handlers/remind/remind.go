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

	timeInput := os.Args[3]
	var deadline time.Time

	if utils.IsDurationFormat(timeInput) {
		deadline, err = utils.ParseDuration(timeInput)
	} else {
		deadline, err = utils.ParseDate(timeInput)
	}

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

	countdown := int(timeToDeadline.Seconds())
	notifyCmd := fmt.Sprintf("sleep %d && notify-send 'Task %d' '%s'", countdown, id, task.Name)
	cmd := exec.Command("bash", "-c", notifyCmd)

	if err = cmd.Start(); err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	fmt.Printf("Reminding task %d in %s\n", id, formatDuration(timeToDeadline))
}

func formatDuration(d time.Duration) string {
	seconds := int(d.Seconds()) % 60
	minutes := int(d.Minutes()) % 60
	hours := int(d.Hours()) % 24
	days := int(d.Hours() / 24)

	parts := []string{}

	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d days", days))
	}

	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d hours", hours))
	}

	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d minutes", minutes))
	}

	if seconds > 0 || len(parts) == 0 { // Aggiungiamo i secondi anche se è zero per evitare una stringa vuota
		parts = append(parts, fmt.Sprintf("%d seconds", seconds))
	}

	if len(parts) > 1 {
		return strings.Join(parts[:len(parts)-1], ", ") + " and " + parts[len(parts)-1]
	}

	return parts[0]
}
