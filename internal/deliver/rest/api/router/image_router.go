package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"golang-project-template/internal/deliver/rest/api/controller"
	"golang-project-template/internal/usecase"
)

func SetupImageRoutes(r chi.Router, db *sql.DB) {
	imageController := controller.ImageController{Usecase: usecase.NewImageUseCase(db)}

	r.Route("/images", func(r chi.Router) {
		r.Post("/", imageController.Create)

		r.Delete("/{key}", imageController.DeleteByIdOrName)

		r.Delete("/", imageController.DeleteAllUnused)

		r.Get("/{key}", imageController.GetByIdOrName)

		r.Get("/", imageController.GetAll)

	})
}
