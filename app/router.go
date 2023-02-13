package app

import (
	"github.com/dzikriqalampacil/cariilmu-course/controller"
	"github.com/dzikriqalampacil/cariilmu-course/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/user", userController.FindAll)
	router.GET("/api/user/:userId", userController.FindById)
	router.POST("/api/user", userController.Create)
	router.PUT("/api/user/:userId", userController.Update)
	router.DELETE("/api/user/:userId", userController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
