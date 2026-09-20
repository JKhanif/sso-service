package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// go run ./cmd/migrator -database-url "postgres://user:pass@localhost:5432/sso?sslmode=disable" -migrations-path ./migrations

func main() {
	var databaseURL, migrationsPath, migrationsTable string

	flag.StringVar(&databaseURL, "database-url", "", "PostgreSQL URL (e.g. postgres://user:pass@host:5432/dbname?sslmode=disable)")
	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "schema_migrations", "name of the migrations table")
	flag.Parse()

	if databaseURL == "" {
		panic("database-url is required")
	}
	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	if migrationsTable != "schema_migrations" {
		sep := "?"
		if strings.Contains(databaseURL, "?") {
			sep = "&"
		}
		databaseURL += sep + "x-migrations-table=" + migrationsTable
	}

	m, err := migrate.New(
		"file://"+migrationsPath,
		databaseURL,
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}

		panic(err)
	}

	fmt.Println("migrations applied")
}
