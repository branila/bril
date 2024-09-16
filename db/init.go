package db

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/branila/bril/types"
)

func Init() {
	raw := getRawDb()

	err := json.Unmarshal(raw, &db)
	if err != nil {
		log.Fatal(err)
	}
}

func getRawDb() []byte {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dbPath := filepath.Join(wd, dbName)

	prepareDb(dbPath)

	content, err := os.ReadFile(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	return content
}

func prepareDb(path string) {
	if !dbExists(path) {
		file := createDb(path)
		defer file.Close()

		setDefault(file)
	}
}

func dbExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func createDb(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func setDefault(file *os.File) {
	jsonData, err := json.Marshal(types.Db{})
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.WriteString(string(jsonData))
	if err != nil {
		log.Fatal(err)
	}
}
