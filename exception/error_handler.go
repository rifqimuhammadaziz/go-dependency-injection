package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rifqimuhammadaziz/go-restful-api/helper"
	"github.com/rifqimuhammadaziz/go-restful-api/model/web"
)

func ErrorHandler(rw http.ResponseWriter, r *http.Request, err interface{}) {
	if notFoundError(rw, r, err) {
		return
	}

	if validationErrors(rw, r, err) {
		return
	}

	internalServerError(rw, r, err)
}

func validationErrors(rw http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest) // set status code(400)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(rw, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(rw http.ResponseWriter, r *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusNotFound) // set status code(404)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(rw, webResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError) // set status code(500)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
