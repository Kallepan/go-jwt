package database

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq" // postgres driver
)

var Instance *sql.DB

func Connect(connectionString string) {
	var err error
	Instance, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
		return
	}

	if err := Instance.Ping(); err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Connected to database!")
}

func Migrate() {
	if Instance == nil {
		log.Fatal("Database is not connected")
	}

	driver, err := postgres.WithInstance(Instance, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)

	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	log.Println("Migrations completed!")
}
