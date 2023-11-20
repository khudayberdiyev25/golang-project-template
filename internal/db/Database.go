package db

import (
	"database/sql"
	"fmt"
	"log"
)

func Setup() *sql.DB {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"db", 5432, "postgres", "root", "app_db",
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
