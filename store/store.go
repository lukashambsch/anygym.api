package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
    "path"
    "runtime"

    "github.com/spf13/viper"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error

    _, filename, _, _ := runtime.Caller(0)
    dir := fmt.Sprintf("%s/..", path.Dir(filename))
    viper.SetConfigName("config")
    viper.AddConfigPath(dir)
    err = viper.ReadInConfig()
    if err != nil {
        fmt.Printf("%#v", err)
    }

	DB, err = Open()
	if err != nil {
		log.Fatal(err)
	}
}

func Open() (*sql.DB, error) {
    env := os.Getenv("GO_ENV")

	connectionInfo := fmt.Sprintf(
		"user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		viper.Get(fmt.Sprintf("datastore.%s.user", env)),
		viper.Get(fmt.Sprintf("datastore.%s.database", env)),
		viper.Get(fmt.Sprintf("datastore.%s.password", env)),
		viper.Get(fmt.Sprintf("datastore.%s.host", env)),
		viper.Get(fmt.Sprintf("datastore.%s.port", env)),
	)
	db, err := sql.Open("postgres", connectionInfo)

	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
