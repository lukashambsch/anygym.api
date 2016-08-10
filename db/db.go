package db

import (
	"database/sql"
	"log"
	"os"
)

func Open() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("GO_DB_URL"))

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
