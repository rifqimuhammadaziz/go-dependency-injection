package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/rifqimuhammadaziz/go-restful-api/app"
	"github.com/rifqimuhammadaziz/go-restful-api/controller"
	"github.com/rifqimuhammadaziz/go-restful-api/helper"
	"github.com/rifqimuhammadaziz/go-restful-api/middleware"
	"github.com/rifqimuhammadaziz/go-restful-api/repository"
	"github.com/rifqimuhammadaziz/go-restful-api/service"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
