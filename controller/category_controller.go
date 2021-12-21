package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(rw http.ResponseWriter, r http.Request, p httprouter.Params)
	Update(rw http.ResponseWriter, r http.Request, p httprouter.Params)
	Delete(rw http.ResponseWriter, r http.Request, p httprouter.Params)
	FindById(rw http.ResponseWriter, r http.Request, p httprouter.Params)
	FindAll(rw http.ResponseWriter, r http.Request, p httprouter.Params)
}
