package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang-project-template/internal/deliver/rest/api/handler"
	"golang-project-template/internal/domain"
	"net/http"
)

type ImageController struct {
	Usecase domain.ImageUseCase
}

func (i *ImageController) Create(writer http.ResponseWriter, request *http.Request) {
	var data domain.ImageRequest
	json.NewDecoder(request.Body).Decode(&data)

	response, err := i.Usecase.Create(&data)
	if err != nil {
		handler.ErrorHandler(writer, request, err)
		return
	}
	render.JSON(writer, request, response)
}

func (i *ImageController) GetAll(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	response, err := i.Usecase.Filter(name)

	if err != nil {
		handler.ErrorHandler(writer, request, err)
		return
	}

	render.JSON(writer, request, response)
}

func (i *ImageController) GetByIdOrName(writer http.ResponseWriter, request *http.Request) {
	key := chi.URLParam(request, "key")
	image, err := i.Usecase.GetOne(key)

	if err != nil {
		handler.ErrorHandler(writer, request, err)
		return
	}

	render.JSON(writer, request, image)
}

func (i *ImageController) DeleteByIdOrName(writer http.ResponseWriter, request *http.Request) {
	key := chi.URLParam(request, "key")

	err := i.Usecase.DeleteByIdOrName(key)

	if err != nil {
		handler.ErrorHandler(writer, request, err)
		return
	}
}

func (i *ImageController) DeleteAllUnused(writer http.ResponseWriter, request *http.Request) {
	response, err := i.Usecase.DeleteAllUnusedOnes()
	if err != nil {
		handler.ErrorHandler(writer, request, err)
		return
	}
	render.JSON(writer, request, response)
}
