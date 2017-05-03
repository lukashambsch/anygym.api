package main

import (
    "errors"
    "fmt"
    "flag"

    _ "github.com/mattes/migrate/driver/postgres"
    "github.com/mattes/migrate/migrate"

    "github.com/lukashambsch/anygym.api/config"
)

func main() {
    ok := false
    errors := []error{errors.New("Pass the -up or -down flag.")}

    up := flag.Bool("up", false, "Set up migrations")
    down := flag.Bool("down", false, "Tear down migrations")

    flag.Parse()

    if *up {
        errors, ok = allUp()
    } else if *down {
        errors, ok = allDown()
    }

    if !ok {
        for _, err := range errors {
            fmt.Printf("%#v", err)
        }
    } else {
        fmt.Println("Migrations run successfully")
    }
}

func allUp() ([]error, bool) {
    errors, ok := migrate.UpSync(fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        config.C.Get("datastore.user"),
        config.C.Get("datastore.password"),
        config.C.Get("datastore.host"),
        config.C.Get("datastore.port"),
        config.C.Get("datastore.database"),
        config.C.Get("datastore.sslmode"),
    ), "./store/migrations")

    return errors, ok
}

func allDown() ([]error, bool) {
    errors, ok := migrate.DownSync(fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        config.C.Get("datastore.user"),
        config.C.Get("datastore.password"),
        config.C.Get("datastore.host"),
        config.C.Get("datastore.port"),
        config.C.Get("datastore.database"),
        config.C.Get("datastore.sslmode"),
    ), "./store/migrations")

    return errors, ok
}
