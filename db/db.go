package db

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/branila/bril/types"
)

const dbName = "bril.json"

var db = types.Db{}

func SyncDb() {
	jsonData, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(wd, dbName)

	err = os.WriteFile(dbPath, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
