package controller

import (
	"net/http"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/web"
	"projectsistemkasir/service"

	"github.com/julienschmidt/httprouter"
)

type ExampleControllerImpl struct {
	ExampleService service.ExampleService
}

func NewExampleController(exampleService service.ExampleService) ExampleController {
	return &ExampleControllerImpl{
		ExampleService: exampleService,
	}
}

func (controller *ExampleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	exampleCreateRequest := web.ExampleCreateRequest{}

	helper.ReadFromRequestBody(request, &exampleCreateRequest)

	exampleResponse := controller.ExampleService.Create(request.Context(), exampleCreateRequest)//membuat datanya
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK luur",
		Data: exampleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}