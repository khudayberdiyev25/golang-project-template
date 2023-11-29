package router

import (
	"database/sql"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang-project-template/internal/domain"
	"log"
	"net/http"
)

func Setup(r chi.Router, db *sql.DB) {
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping"))
	})

	SetupImageRoutes(r, db)
	SetupContainerRoutes(r, db)
	SetupVolumeRoutes(r, db)
	SetupNetworkRoutes(r, db)

	r.Post("/example", func(w http.ResponseWriter, r *http.Request) {
		var data domain.ImageRequest

		// Decode JSON request body into the struct
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}

		// Do something with the data (e.g., process it, validate it)
		// For simplicity, just print it in this example
		log.Printf("%+v", data)
		render.JSON(w, r, data)
	})
}
