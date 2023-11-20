package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"golang-project-template/internal/bootstrap"
	"golang-project-template/internal/deliver/rest/api/router"
	"log"
	"net/http"
)

func main() {
	bootstrap.LoadEnv()
	db := bootstrap.SetupDB()
	r := chi.NewRouter()
	router.Setup(r, db)
	err := http.ListenAndServe(":5005", r)
	if err != nil {
		log.Fatal(err)
	}
}
