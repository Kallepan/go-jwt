package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgres driver
)

var Instance *sql.DB

func Connect(connectionString string) error {
	Instance, err := sql.Open("postgres", connectionString)

	if err != nil {
		return err
	}

	err = Instance.Ping()

	if err != nil {
		return err
	}

	log.Println("Connected to database!")
	return nil
}
