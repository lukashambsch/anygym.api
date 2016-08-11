package store

import (
	"database/sql"
	"log"
	"os"
)

var DB *sql.DB

func init() {
	var err error

	DB, err = Open()
	if err != nil {
		log.Fatal(err)
	}
}

func Open() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("GO_DB_URL"))

	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
