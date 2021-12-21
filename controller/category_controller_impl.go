package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rifqimuhammadaziz/go-restful-api/helper"
	"github.com/rifqimuhammadaziz/go-restful-api/model/web"
	"github.com/rifqimuhammadaziz/go-restful-api/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService // service.CategoryService is interface, no need to set pointer
}

func (controller *CategoryControllerImpl) Create(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CategoryControllerImpl) Update(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryUpdateRequest)

	categoryId := p.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CategoryControllerImpl) Delete(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId") // get id
	id, err := strconv.Atoi(categoryId)  // convert to string
	helper.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CategoryControllerImpl) FindById(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryId := p.ByName("categoryId") // get id
	id, err := strconv.Atoi(categoryId)  // convert to string
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(rw, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.WriteToResponseBody(rw, webResponse)
}
