package controller

import (
	"encoding/json"
	"github.com/go-chi/render"
	"golang-project-template/internal/deliver/rest/api/handler"
	"golang-project-template/internal/domain"
	"net/http"
)

type ContainerController struct {
	UseCase domain.ContainerUseCase
}

func (c *ContainerController) Create(writer http.ResponseWriter, request *http.Request) {
	var data domain.ContainerRequest
	json.NewDecoder(request.Body).Decode(&data)

	response, err := c.UseCase.Create(&data)
	if err != nil {
		handler.ErrorHandler(writer, request, err)
		return
	}

	render.JSON(writer, request, response)
}
