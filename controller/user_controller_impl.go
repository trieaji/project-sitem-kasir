package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/web"
	"projectsistemkasir/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}

	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)//membuat datanya

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK luur",
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.LoginUser{}

	helper.ReadFromRequestBody(request, &userCreateRequest)

	userResponse := controller.UserService.Login(request.Context(), userCreateRequest)
	fmt.Println("===userResponselogin===")
	fmt.Println(userResponse)

	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

	// writer.Header().Add("Content-Type", "application/json")
	// encoder := json.NewEncoder(writer)
	// err = encoder.Encode(webResponse)
	// helper.PanicIfError(err)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	userUpdateRequest := web.UserUpdateRequest{}
	err := decoder.Decode(&userUpdateRequest)
	helper.PanicIfError(err)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: userResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), id)
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *UserControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), id)
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: userResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *UserControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userresponses := controller.UserService.FindAll(request.Context())
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: userresponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *UserControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.LogoutUser{}

	helper.ReadFromRequestBody(request, &userCreateRequest)

	controller.UserService.Logut(request.Context(), userCreateRequest)

	webResponse := web.WebResponseLogout {
		Data: nil,
	}

	// helper.WriteToResponseBody(writer, webResponse)
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}