package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"golang-project-template/internal/deliver/rest/api/controller"
	"golang-project-template/internal/domain"
	"golang-project-template/internal/usecase"
)

func SetupContainerRoutes(r chi.Router, db *sql.DB) {
	containerController := controller.ContainerController{UseCase: usecase.NewContainerUseCase(db, &domain.ContainerMapper{})}
	r.Route("/containers", func(r chi.Router) {
		r.Post("/", containerController.Create)
	})
}
