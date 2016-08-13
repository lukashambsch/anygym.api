package store

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", os.Getenv("DATABASE_CONFIG"))

	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
