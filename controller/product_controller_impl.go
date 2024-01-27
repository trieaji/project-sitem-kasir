package controller

import (
	"encoding/json"
	"net/http"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/web"
	"projectsistemkasir/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productCreateRequest := web.ProductCreateRequest{}

	helper.ReadFromRequestBody(request, &productCreateRequest)

	productResponse := controller.ProductService.Create(request.Context(), productCreateRequest)//membuat datanya
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK luur",
		Data: productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	productUpdateRequest := web.ProductUpdateRequest{}
	err := decoder.Decode(&productUpdateRequest)
	helper.PanicIfError(err)

	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productUpdateRequest.Id = id

	productResponse := controller.ProductService.Update(request.Context(), productUpdateRequest)
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: productResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: productResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productresponses := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: productresponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}