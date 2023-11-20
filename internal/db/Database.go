package db

import (
	"database/sql"
	"fmt"
)

func Setup() *sql.DB {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"db", 5432, "postgres", "root", "app_db",
	)

	db, _ := sql.Open("postgres", connString)

	return db
}
