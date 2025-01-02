package db

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/branila/bril/types"
)

const dbName = "bril.json"

var db = types.Db{}

func SyncDb() error {
	jsonData, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling data: %e", err)
	}

	err = os.WriteFile(dbPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}

	return nil
}

func Reset() error {
	db = types.Db{}
	return SyncDb()
}
