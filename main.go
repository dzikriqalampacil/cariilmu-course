package main

import (
	"net/http"

	"github.com/dzikriqalampacil/cariilmu-course/app"
	"github.com/dzikriqalampacil/cariilmu-course/controller"
	"github.com/dzikriqalampacil/cariilmu-course/helper"
	"github.com/dzikriqalampacil/cariilmu-course/repository"
	"github.com/dzikriqalampacil/cariilmu-course/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := app.NewRouter(userController)

	server := http.Server{
		Addr:    "localhost:8001",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
