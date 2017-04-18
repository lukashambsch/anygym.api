package store

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

    "github.com/lukashambsch/anygym.api/config"
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
    var err error
    var db *sql.DB

	connectionInfo := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.C.Get("datastore.user"),
		config.C.Get("datastore.database"),
		config.C.Get("datastore.password"),
		config.C.Get("datastore.host"),
		config.C.Get("datastore.port"),
	)

    db, err = sql.Open("postgres", connectionInfo)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
