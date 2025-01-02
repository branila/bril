// init.go
package db

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/branila/bril/types"
)

var dbPath string

func Init() {
	dbPath = getDbPath()

	fmt.Println("dbPath:", dbPath)

	prepareDb(dbPath)

	content, err := os.ReadFile(dbPath)
	if err != nil {
		log.Fatal("Failed to read database:", err)
	}

	err = json.Unmarshal(content, &db)
	if err != nil {
		log.Fatal("Failed to parse database:", err)
	}
}

func getDbPath() string {
	userDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Failed to get user config directory:", err)
	}

	brilDir := filepath.Join(userDir, "bril")
	if err := os.MkdirAll(brilDir, 0755); err != nil {
		log.Fatal("Failed to create bril directory:", err)
	}

	return filepath.Join(brilDir, dbName)
}

func prepareDb(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		db := types.Db{}
		jsonData, err := json.Marshal(db)
		if err != nil {
			log.Fatal("Failed to create empty database:", err)
		}

		if err := os.WriteFile(path, jsonData, 0644); err != nil {
			log.Fatal("Failed to write empty database:", err)
		}
	}
}
