package common

import (
	"database/sql"
	"log"
	"os"
)

func GetPathHere() string {
	dir, err := os.Getwd()
	if err != nil {
		return "." + Separator
	}
	return dir
}

func ConnectDB(source string) {
	var err error
	DB, err = sql.Open("sqlite3", source)
	if err != nil {
		log.Fatal(err)
	}
}
