package resetter

import (
	"fmt"
	"strings"

	"github.com/branila/bril/db"
)

func Reset() {
	fmt.Print("Are you sure you want to reset the program? (Y/n): ")

	var response string
	fmt.Scanln(&response)

	if strings.ToLower(response) != "y" {
		fmt.Println("\nReset aborted.")
		return
	}

	if err := db.Reset(); err != nil {
		fmt.Println("\nReset failed:", err)
		return
	}

	fmt.Println("\nProgram successfully reset.")
}
