package tags

import (
	"fmt"
	"os"
)

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

	case "edit":
		Edit()

	default:
		fmt.Println("Usage: bril tag <command> [args]")
	}
}
