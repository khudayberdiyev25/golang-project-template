package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"golang-project-template/internal/db"
	"golang-project-template/internal/deliver/rest/api/router"
	"net/http"
)

func main() {
	db := db.Setup()
	r := chi.NewRouter()
	router.Setup(r, db)
	http.ListenAndServe(":5005", r)
}
