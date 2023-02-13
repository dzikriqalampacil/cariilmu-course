package controller

import (
	"net/http"
	"strconv"

	"github.com/dzikriqalampacil/cariilmu-course/helper"
	"github.com/dzikriqalampacil/cariilmu-course/model/web"
	"github.com/dzikriqalampacil/cariilmu-course/service"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(UserService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: UserService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserCreateRequest := web.UserCreateRequest{}
	helper.ReadFromRequestBody(request, &UserCreateRequest)

	UserResponse := controller.UserService.Create(request.Context(), UserCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   UserResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &UserUpdateRequest)

	UserId := params.ByName("UserId")
	id, err := strconv.ParseInt(UserId, 10, 64)
	helper.PanicIfError(err)

	UserUpdateRequest.Id = id

	UserResponse := controller.UserService.Update(request.Context(), UserUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   UserResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserId := params.ByName("UserId")
	id, err := strconv.ParseInt(UserId, 10, 64)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserId := params.ByName("UserId")
	id, err := strconv.ParseInt(UserId, 10, 64)
	helper.PanicIfError(err)

	UserResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   UserResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	UserResponses := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   UserResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
