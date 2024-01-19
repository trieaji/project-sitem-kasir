package service

import (
	"context"
	"database/sql"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/domain"
	"projectsistemkasir/model/web"
	"projectsistemkasir/repository"

	"github.com/go-playground/validator/v10"
)

type ExampleServiceImpl struct {
	ExampleRepository repository.ExampleRepository
	DB *sql.DB
	Validate validator.Validate
}

func NewExampleService(exampleRepository repository.ExampleRepository, DB *sql.DB, validate *validator.Validate) ExampleService {
	return &ExampleServiceImpl{
		ExampleRepository: exampleRepository,
		DB: DB,
		Validate: *validate,
	}
}

func (service *ExampleServiceImpl) Create(ctx context.Context, request web.ExampleCreateRequest) web.ExampleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	example := domain.Examples {
		Name: request.Name,
		Jurusan: request.Jurusan,
	}

	example = service.ExampleRepository.Save(ctx,tx,example)

	return helper.ToExampleResponse(example)
	
}