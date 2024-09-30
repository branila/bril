package main

import (
	"fmt"
	"os"

	"github.com/branila/bril/db"
	"github.com/branila/bril/handlers/adder"
)

func main() {
	db.Init()

	if len(os.Args) < 2 {
		fmt.Println("Usage: bril <command> [args]")
		return
	}

	switch os.Args[1] {
	case "add":
		adder.Add()

	default:
		fmt.Println("Unknown command")
	}
}
