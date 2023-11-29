package handler

import (
	"github.com/go-chi/render"
	"golang-project-template/internal/domain"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err error) {
	var errorResponse domain.ErrorResponse
	switch err.(type) {
	case *domain.ErrImageNotFound:
		errorResponse = domain.ErrorResponse{
			Err:    err.Error(),
			Status: err.(*domain.ErrImageNotFound).Status,
		}
	case *domain.ErrImageInUse:
		errorResponse = domain.ErrorResponse{
			Err:    err.Error(),
			Status: err.(*domain.ErrImageInUse).Status,
		}
	case *domain.ErrImageDoesNotExist:
		errorResponse = domain.ErrorResponse{
			Err:    err.Error(),
			Status: err.(*domain.ErrImageDoesNotExist).Status,
		}
	default:
		errorResponse = domain.ErrorResponse{
			Err:    err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	writer.WriteHeader(errorResponse.Status)
	render.JSON(writer, request, errorResponse)
}
