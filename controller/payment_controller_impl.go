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

type PaymentControllerImpl struct {
	PaymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) PaymentController {
	return &PaymentControllerImpl{
		PaymentService: paymentService,
	}
}

func (controller *PaymentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentCreateRequest := web.PaymentCreateRequest{}

	helper.ReadFromRequestBody(request, &paymentCreateRequest)

	paymentResponse := controller.PaymentService.Create(request.Context(), paymentCreateRequest)//membuat datanya

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK luur",
		Data: paymentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}


func (controller *PaymentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	paymentUpdateRequest := web.PaymentUpdateRequest{}
	err := decoder.Decode(&paymentUpdateRequest)
	helper.PanicIfError(err)

	paymentId := params.ByName("paymentId")
	id, err := strconv.Atoi(paymentId)
	helper.PanicIfError(err)

	paymentUpdateRequest.Id = id

	paymentResponse := controller.PaymentService.Update(request.Context(), paymentUpdateRequest)
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: paymentResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}


func (controller *PaymentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentId := params.ByName("paymentId")
	id, err := strconv.Atoi(paymentId)
	helper.PanicIfError(err)

	controller.PaymentService.Delete(request.Context(), id)
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *PaymentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentId := params.ByName("paymentId")
	id, err := strconv.Atoi(paymentId)
	helper.PanicIfError(err)

	paymentResponse := controller.PaymentService.FindById(request.Context(), id)
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: paymentResponse,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	helper.PanicIfError(err)
}

func (controller *PaymentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentresponses := controller.PaymentService.FindAll(request.Context())
	webResponse := web.WebResponse {
		Code: 200,
		Status: "OK luur",
		Data: paymentresponses,
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	helper.PanicIfError(err)
}