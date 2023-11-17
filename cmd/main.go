package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"db", 5432, "postgres", "root", "app_db",
	)

	db, _ := sql.Open("postgres", connString)
	//
	defer db.Close()

	http.ListenAndServe(":5005", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		err := db.Ping()
		if err != nil {
			log.Fatal(err)
		} else {
			var id int
			err := db.QueryRow(`insert into images(name) values ($1) returning id`, "simple-docker-app").Scan(&id)
			if err != nil {
				writer.Write([]byte(err.Error()))
			} else {
				fmt.Println(id)
				writer.Write([]byte(string(rune(id))))
			}
		}

	}))
}
