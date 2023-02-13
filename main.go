package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dzikriqalampacil/cariilmu-course/app"
	"github.com/dzikriqalampacil/cariilmu-course/controller"
	"github.com/dzikriqalampacil/cariilmu-course/helper"
	"github.com/dzikriqalampacil/cariilmu-course/repository"
	"github.com/dzikriqalampacil/cariilmu-course/service"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Some error occured. Err: %s", errEnv)
	}

	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	db := app.NewDB(host, port, user, password, dbname)

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
