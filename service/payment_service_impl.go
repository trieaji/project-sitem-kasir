package service

import (
	"context"
	"database/sql"
	"projectsistemkasir/exception"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/domain"
	"projectsistemkasir/model/web"
	"projectsistemkasir/repository"

	"github.com/go-playground/validator/v10"
)

type PaymentServiceImpl struct {
	PaymentRepository repository.PaymentRepository
	DB *sql.DB
	Validate validator.Validate
}

func NewPaymentService(paymentRepository repository.PaymentRepository, DB *sql.DB, validate *validator.Validate) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository: paymentRepository,
		DB: DB,
		Validate: *validate,
	}
}

func (service *PaymentServiceImpl) Create(ctx context.Context, request web.PaymentCreateRequest) web.PaymentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payment := domain.Payments {
		Name: request.Name,
		Type: request.Type,
	}

	payment = service.PaymentRepository.Save(ctx,tx,payment)

	return helper.ToPaymentResponse(payment)
}

func (service *PaymentServiceImpl) Update(ctx context.Context, request web.PaymentUpdateRequest) web.PaymentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	pay, err := service.PaymentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// payment := domain.Payments {
	// 	pay.Name : request.Name,
	// 	pay.Type : request.Name,
	// }

	

	pay.Name = request.Name
	pay.Type = request.Type

	pay = service.PaymentRepository.Update(ctx,tx,pay)

	return helper.ToPaymentResponse(pay)
}

func (service *PaymentServiceImpl) Delete(ctx context.Context, paymentId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payment, err := service.PaymentRepository.FindById(ctx, tx, paymentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.PaymentRepository.Delete(ctx, tx, payment)
}

func (service *PaymentServiceImpl) FindById(ctx context.Context, paymentId int) web.PaymentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payment, err := service.PaymentRepository.FindById(ctx, tx, paymentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPaymentResponse(payment)
}

func (service *PaymentServiceImpl) FindAll(ctx context.Context) []web.PaymentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	payments := service.PaymentRepository.FindAll(ctx, tx)

	return helper.ToPaymentResponses(payments)
}