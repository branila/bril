package tags

import (
	"fmt"
	"os"
)

func List() {
	fmt.Println("Coming soon")
}

func Delete() {
	fmt.Println("Coming soon")
}

func Handle() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: bril tag <command> [args]")
		return
	}

	switch os.Args[2] {
	case "add":
		Add()
	case "ls", "list":
		List()
	case "rm", "delete":
		Delete()
	default:
		fmt.Println("Unknown command")
	}
}
